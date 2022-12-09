package main

import (
	"GoGit-Integration/pkg/config"
	"GoGit-Integration/pkg/gitapi"
	"GoGit-Integration/pkg/goGit"

	logr "github.com/sirupsen/logrus"
)

func main() {
	logr.Infoln("Grabbing configuration...")
	config, err := config.GetConfig()
	if err != nil {
		logr.Panic(err)
	}
	logr.SetLevel(logr.Level(config.LogLevel))

	names := gitapi.GetList(config)

	goGit.UpdateLocalCopies(names, config)
}
