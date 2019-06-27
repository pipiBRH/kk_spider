package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		GovDatasetURL string              `yaml:"govDatasetURL"`
		Elasticsearch ElasticsearchConfig `yaml:"elasticsearch"`
	}

	ElasticsearchConfig struct {
		Host  string `yaml:"host"`
		Port  int    `yaml:"port"`
		Index string `yaml:"index"`
		Type  string `yaml:"type"`
	}
)

func NewConfig(path string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := new(Config)
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
