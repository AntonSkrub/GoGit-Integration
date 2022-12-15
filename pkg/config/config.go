package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var instance *Config
var configPath = "./config/config.yml"

type Config struct {
	OrgaName  string `yaml:"OrgaName"`
	OrgaToken string `yaml:"OrgaToken"`

	UserName  string `yaml:"UserName"`
	UserToken string `yaml:"UserToken"`

	OutputPath string `yaml:"OutputPath"`

	ListReferences   bool `yaml:"ListReferences"`
	LogLatestCommits bool `yaml:"LogCommits"`
	LogLevel         int  `yaml:"LogLevel"`
}

func GetConfig() (*Config, error) {
	instance = &Config{}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return instance, fmt.Errorf("[config] Couldn't find config file: %s", err.Error())
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return instance, fmt.Errorf("[config] Could not read config file: %s", err.Error())
	}

	err = yaml.Unmarshal(file, instance)
	if err != nil {
		return instance, fmt.Errorf("[config] Error parsing the config: %s", err.Error())
	}

	return instance, nil
}
