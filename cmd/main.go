package main

import (
	"GoGit-Integration/pkg/config"
	"GoGit-Integration/pkg/gitapi"
	"GoGit-Integration/pkg/goGit"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"
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

	logr.Info("Creating a stop channel for the cron job ...")

	UpdateInterval := cron.New()
	UpdateInterval.AddFunc("*/3 * * * *", func() {
		goGit.UpdateLocalCopies(names, config)
	})
	go UpdateInterval.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	logr.Info("Received an interrupt signal, stopping the cron jobs ...")

	UpdateInterval.Stop()
	logr.Info("closed the stop channel, exiting ...")
	os.Exit(0)
}
