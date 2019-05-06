package helper

type TaskConfig struct {
	Inputs  []map[string]string `yaml:"inputs"`
	Outputs []map[string]string `yaml:"outputs"`
	Params  map[string]string   `yaml:"params"`
}
