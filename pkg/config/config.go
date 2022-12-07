package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var instance *Config
var configPath = "./config/config.yml"

type Config struct {
	OrgaToken string `yaml:"OrgaToken"`
}

func GetConfig() (*Config, error) {
	instance = &Config{}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return instance, fmt.Errorf("config file not found: %s", err.Error())
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return instance, fmt.Errorf("error opening config file: %s", err.Error())
	}

	err = yaml.Unmarshal(file, instance)
	if err != nil {
		return instance, fmt.Errorf("error parsing config file: %s", err.Error())
	}

	return instance, nil
}
