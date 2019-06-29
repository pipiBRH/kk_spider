package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type (
	ConfigData struct {
		GovDatasetURL string              `yaml:"govDatasetURL"`
		Elasticsearch ElasticsearchConfig `yaml:"elasticsearch"`
	}

	ElasticsearchConfig struct {
		Host  string            `yaml:"host"`
		Port  int               `yaml:"port"`
		Index map[string]string `yaml:"index"`
	}
)

var Config ConfigData

func NewConfig(path string) error {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		return err
	}

	return nil
}
