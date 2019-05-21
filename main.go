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
	yaml "gopkg.in/yaml.v1"

	"github.com/ankeesler/flexec/helper"
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

	flexecConfig := helper.NewFlexecConfig()
	if err := helper.ReadFlexecConfig(flexecConfig); err != nil {
		return errors.Wrap(err, "load flexexec config")
	}
	defer func() {
		if err := helper.WriteFlexecConfig(flexecConfig); err != nil {
			log.Printf("note: error: write flexec config: %s", err.Error())
		}
	}()

	taskName := os.Args[1]
	taskConfigPath := filepath.Join(
		os.Getenv("HOME"),
		"workspace",
		"credhub-ci",
		"tasks",
		taskName,
		"task.yml",
	)
	taskConfigData, err := getTaskConfigData(taskConfigPath)
	if err != nil {
		return errors.Wrap(err, "get task config data")
	}

	var taskConfig helper.TaskConfig
	if err := yaml.Unmarshal(taskConfigData, &taskConfig); err != nil {
		return errors.Wrap(err, "unmarshal task config yaml")
	}

	cmdBuilder := helper.NewCmdBuilder(taskConfigPath)
	if err := resolveInputs(&taskConfig, cmdBuilder.OnInput); err != nil {
		return errors.Wrap(err, "resolve inputs")
	}
	if err := resolveOutputs(&taskConfig, cmdBuilder.OnOutput); err != nil {
		return errors.Wrap(err, "resolve outputs")
	}
	if err := resolveParams(
		&taskConfig,
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

func resolveInputs(config *helper.TaskConfig, handler func(name, path string)) error {
	for _, inputMap := range config.Inputs {
		input, ok := inputMap["name"]
		if !ok {
			return fmt.Errorf("input map does not contain name: %s", inputMap)
		}

		path, ok := resolveRepo(input)
		if !ok {
			path = fmt.Sprintf("/tmp/%s", input)
			if err := os.MkdirAll(path, 0700); err != nil {
				return errors.Wrap(err, fmt.Sprintf("mkdir (%s)", path))
			}
			log.Printf("cannot resolve repo for input %s", input)
		}
		log.Printf("resolved input %s to path %s", input, path)
		handler(input, path)
	}
	return nil
}

func resolveOutputs(config *helper.TaskConfig, handler func(name, path string)) error {
	for _, outputMap := range config.Outputs {
		output, ok := outputMap["name"]
		if !ok {
			return fmt.Errorf("output map does not contain name: %s", outputMap)
		}

		path := fmt.Sprintf("/tmp/%s", output)
		log.Printf("resolved output %s to path %s", output, path)
		handler(output, path)
	}
	return nil
}

func resolveParams(
	config *helper.TaskConfig,
	defaults map[string]string,
	handler func(paramKey, paramValue string),
) error {
	for paramKey, paramValue := range config.Params {
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
