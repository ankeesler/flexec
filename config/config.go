package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v1"
)

type Config struct {
	ParamDefaults map[string]string `yaml:"param_defaults"`
}

func Read(storeDir string) (*Config, error) {
	c := &Config{
		ParamDefaults: make(map[string]string),
	}

	configFile := flexecConfigFile(storeDir)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return c, nil
	}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, errors.Wrap(err, "read file")
	}

	if err := yaml.Unmarshal(data, c); err != nil {
		return nil, errors.Wrap(err, "unmarshal yaml")
	}

	return c, nil
}

func Write(storeDir string, c *Config) error {
	configFile := flexecConfigFile(storeDir)

	data, err := yaml.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "marshal yaml")
	}

	if err := ioutil.WriteFile(configFile, data, 0600); err != nil {
		return errors.Wrap(err, "write file")
	}

	return nil
}

func flexecConfigFile(storeDir string) string {
	return filepath.Join(storeDir, ".flexec")
}
