package helper

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type CmdBuilder struct {
	cmd *exec.Cmd
}

func NewCmdBuilder(taskConfigPath string) *CmdBuilder {
	cmd := exec.Command(
		"fly",
		"--target",
		"credhub",
		"execute",
		"--config",
		taskConfigPath,
	)
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", "HOME", os.Getenv("HOME")))

	return &CmdBuilder{cmd: cmd}
}

func (cb *CmdBuilder) OnInput(name, path string) {
	inputArg := fmt.Sprintf("--input=%s=%s", name, path)
	cb.cmd.Args = append(cb.cmd.Args, inputArg)
}

func (cb *CmdBuilder) OnOutput(name, path string) {
	outputArg := fmt.Sprintf("--output=%s=%s", name, path)
	cb.cmd.Args = append(cb.cmd.Args, outputArg)
}

func (cb *CmdBuilder) OnParam(paramKey, paramValue string) {
	env := fmt.Sprintf("%s=%s", paramKey, paramValue)
	cb.cmd.Env = append(cb.cmd.Env, env)
}

func (cb *CmdBuilder) Build() *exec.Cmd {
	return cb.cmd
}

func (cb *CmdBuilder) String() string {
	cmd := cb.Build()
	s := make([]string, 10)

	for _, env := range cmd.Env {
		s = append(s, env)
		s = append(s, "\n  ")
	}

	for _, arg := range cmd.Args {
		s = append(s, arg)
		s = append(s, "\n  ")
	}

	return strings.Join(s, "")
}
