package helper

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v1"
)

type FlexecConfig struct {
	ParamDefaults map[string]string `yaml:"param_defaults"`
}

func NewFlexecConfig() *FlexecConfig {
	return &FlexecConfig{
		ParamDefaults: make(map[string]string),
	}
}

func ReadFlexecConfig(config *FlexecConfig) error {
	configFile := flexecConfigFile()
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil
	}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return errors.Wrap(err, "read flexec config")
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return errors.Wrap(err, "unmarshal flexec config")
	}

	return nil
}

func WriteFlexecConfig(config *FlexecConfig) error {
	configFile := flexecConfigFile()

	data, err := yaml.Marshal(config)
	if err != nil {
		return errors.Wrap(err, "marshal flexec config")
	}

	if err := ioutil.WriteFile(configFile, data, 0600); err != nil {
		return errors.Wrap(err, "unmarshal flexec config")
	}

	return nil
}

func flexecConfigFile() string {
	return filepath.Join(
		os.Getenv("HOME"),
		".flexec",
	)
}
