package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/AntonSkrub/GoGit-Integration/pkg/config"
	"github.com/AntonSkrub/GoGit-Integration/pkg/gitapi"
	"github.com/AntonSkrub/GoGit-Integration/pkg/gogit"
	"github.com/robfig/cron/v3"
	logr "github.com/sirupsen/logrus"
)

func main() {
	logr.Infoln("Grabbing configuration...")
	config, err := config.GetConfig()
	if err != nil {
		logr.Fatalf("Failed to get configuration: %v", err)
	}
	logr.SetLevel(logr.Level(config.LogLevel))

	names := gitapi.GetList(config)
	gogit.UpdateLocalCopies(names, config)

	UpdateInterval := cron.New()
	UpdateInterval.AddFunc("*/3 * * * *", func() {
		gogit.UpdateLocalCopies(names, config)
	})
	go UpdateInterval.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	logr.Info("Received an interrupt signal, stopping the cron jobs ...")

	UpdateInterval.Stop()
	logr.Info("Cron jobs stopped, exiting ...")
	os.Exit(0)
}
