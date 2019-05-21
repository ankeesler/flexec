package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/ankeesler/flexec/config"
	"github.com/ankeesler/flexec/helper"
	"github.com/ankeesler/flexec/task"
	"github.com/concourse/concourse/atc"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}

func run() error {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <task-name>\n", os.Args[0])
		os.Exit(1)
	}

	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", filepath.Base(os.Args[0])))

	flexecConfig, err := config.Read(os.Getenv("HOME"))
	if err != nil {
		return errors.Wrap(err, "load flexexec config")
	}
	defer func() {
		if err := config.Write(os.Getenv("HOME"), flexecConfig); err != nil {
			log.Printf("note: error: write flexec config: %s", err.Error())
		}
	}()

	taskName := os.Args[1]
	taskPath := filepath.Join(
		os.Getenv("HOME"),
		"workspace",
		"credhub-ci",
		"tasks",
		taskName,
		"task.yml",
	)
	task, err := task.Read(taskPath)
	if err != nil {
		return errors.Wrap(err, "read task")
	}

	cmdBuilder := helper.NewCmdBuilder(taskPath)
	if err := resolveInputs(task, cmdBuilder.OnInput); err != nil {
		return errors.Wrap(err, "resolve inputs")
	}
	if err := resolveOutputs(task, cmdBuilder.OnOutput); err != nil {
		return errors.Wrap(err, "resolve outputs")
	}
	if err := resolveParams(
		task,
		flexecConfig.ParamDefaults,
		cmdBuilder.OnParam,
	); err != nil {
		return errors.Wrap(err, "resolve params")
	}

	cmd := cmdBuilder.Build()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("running command: \n  %s", cmdBuilder.String())
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "run command")
	}

	return nil
}

func getTaskConfigData(taskConfigPath string) ([]byte, error) {
	log.Printf("reading task path %s", taskConfigPath)

	taskConfigData, err := ioutil.ReadFile(taskConfigPath)
	if err != nil {
		return nil, errors.Wrap(err, "read task config file")
	}

	return taskConfigData, nil
}

func resolveInputs(task *atc.TaskConfig, handler func(name, path string)) error {
	for _, input := range task.Inputs {
		path, ok := resolveRepo(input.Name)
		if !ok {
			path = fmt.Sprintf("/tmp/%s", input.Name)
			if err := os.MkdirAll(path, 0700); err != nil {
				return errors.Wrap(err, fmt.Sprintf("mkdir (%s)", path))
			}
			log.Printf("cannot resolve repo for input %s", input.Name)
		}
		log.Printf("resolved input %s to path %s", input.Name, path)
		handler(input.Name, path)
	}
	return nil
}

func resolveOutputs(task *atc.TaskConfig, handler func(name, path string)) error {
	for _, output := range task.Outputs {
		path := fmt.Sprintf("/tmp/%s", output.Name)
		log.Printf("resolved output %s to path %s", output.Name, path)
		handler(output.Name, path)
	}
	return nil
}

func resolveParams(
	task *atc.TaskConfig,
	defaults map[string]string,
	handler func(paramKey, paramValue string),
) error {
	for paramKey, paramValue := range task.Params {
		if paramValue == "" {
			fmt.Printf(
				"enter value for param %s (default '%s'): ",
				paramKey,
				defaults[paramKey],
			)
			text, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				return errors.Wrap(err, "read string from stdin")
			}
			paramValue = strings.TrimSpace(text)
			if len(paramValue) == 0 {
				paramValue = defaults[paramKey]
			}
			defaults[paramKey] = paramValue
		}

		log.Printf("setting param: %s => %s", paramKey, paramValue)
		handler(paramKey, paramValue)
	}
	return nil
}

func resolveRepo(name string) (string, bool) {
	repo := filepath.Join(
		os.Getenv("HOME"),
		"workspace",
		name,
	)
	if _, err := os.Stat(repo); err == nil {
		return repo, true
	}

	cmd := exec.Command(
		"find",
		filepath.Join(os.Getenv("HOME"), "go"),
		"-name",
		name,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("note: failed to run find command (output = %s)", string(output))
		return "", false
	}

	outputString := string(output)
	if strings.Count(outputString, "\n") == 1 {
		return strings.TrimSpace(outputString), true
	}

	return "", false
}
