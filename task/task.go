package task

import (
	"io/ioutil"
	"log"

	"github.com/concourse/concourse/atc"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v1"
)

func Read(path string) (*atc.TaskConfig, error) {
	log.Printf("reading task path %s", path)

	taskConfigData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "read file")
	}

	task := new(atc.TaskConfig)
	if err := yaml.Unmarshal(taskConfigData, &task); err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}

	return task, nil
}
