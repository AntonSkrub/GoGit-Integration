package config

import (
	"os"
	"path/filepath"

	logr "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var instance *Config
var configPath = "./config"

type Config struct {
	OrgaName  string `yaml:"OrgaName"`
	OrgaToken string `yaml:"OrgaToken"`

	UserName  string `yaml:"UserName"`
	UserToken string `yaml:"UserToken"`

	OutputPath string `yaml:"OutputPath"`

	ListReferences bool `yaml:"ListReferences"`
	LogCommits     bool `yaml:"LogCommits"`
	LogLevel       int  `yaml:"LogLevel"`
}

func GetConfig() *Config {
	if instance == nil {
		err := initConfig()
		if err != nil {
			logr.Fatalf("[config] Error initializing the config: %s", err.Error())
		}
	}

	return instance
}

func initConfig() error {
	instance = &Config{}

	if _, err := os.Stat(filepath.Join(configPath, "config.yml")); err != nil {
		err = createConfig()
		if err != nil {
			return err
		}
	}

	file, err := os.Open(filepath.Join(configPath, "config.yml"))
	if err != nil {
		return err
	}
	defer file.Close()

	err = yaml.NewDecoder(file).Decode(instance)
	if err != nil {
		return err
	}

	return nil
}

func createConfig() error {
	config := &Config{
		OrgaName: "Default Orga",
		OrgaToken: "",
		UserName: "Default User",
		UserToken: "",
		OutputPath: "../Repo-Backups/",
		ListReferences: true,
		LogCommits: false,
		LogLevel: 6,
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(configPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	err = os.WriteFile(filepath.Join(configPath, "config.yml"), data, 0600)
	if err != nil {
		return err
	}
	logr.Info("[config] created default configuration, exiting...")
	os.Exit(0)
	return nil
}
