package handlers

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type HandlersConfig struct {
	Calls []CallConfig `yaml:"calls"`
}

type CallConfig struct {
	Method  string            `yaml:"method"`
	Params  map[string]string `yaml:"params"`
	Backend []BackendConfig   `yaml:"backend"`
}

type BackendConfig struct {
	ResponseBody    string   `yaml:"response_body"`
	Allow           []string `yaml:"allow"`
	RequestTemplate string   `yaml:"request_template"`
}

func LoadConfig(path string) (*HandlersConfig, error) {
	// Read the YAML file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal the YAML data into HandlersConfig
	var config HandlersConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
