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
	Accounts map[string]Account `yaml:"Accounts"`

	OutputPath     string `yaml:"OutputPath"`
	UpdateInterval string `yaml:"UpdateInterval"`

	ListReferences bool   `yaml:"ListReferences"`
	LogLevel       uint32 `yaml:"LogLevel"`
}

type Account struct {
	Name         string `yaml:"Name"`
	Token        string `yaml:"Token"`
	Option       string `yaml:"Option"`
	BackupRepos  bool   `yaml:"BackupRepos"`
	ValidateName bool   `yaml:"ValidateName"` // Whether the User-/OrgaName has to be contained in the "full_name" of the repository
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
	logr.Info("[config] Creating default configuration ...")

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
		Accounts: map[string]Account{
			"1st": {
				Name:         "GitHub-Username",
				Token:        "Github-Access-Token",
				Option:       "all",
				BackupRepos:  true,
				ValidateName: false,
			},
		},
		OutputPath:     "../Repo-Backups/",
		UpdateInterval: "0 */12 * * * *",

		ListReferences: true,
		LogLevel:       6,
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

func SetConfigPath(path string) {
	configPath = path
}
