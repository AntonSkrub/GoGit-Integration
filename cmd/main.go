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
	config := config.GetConfig()
	logr.SetLevel(logr.Level(config.LogLevel))

	repoNames := []gitapi.Repository{}
	for _, account := range config.Accounts {
		if !account.BackupRepos {
			logr.Infof("[main] Skipping account %v because it's not set to backup repositories", account.Name)
			continue
		}

		logr.Printf("[API] Found account: %v", account)
		repoNames = gitapi.GetRepoList(&account)
		logr.Info("[main] Found ", len(repoNames), " repositories on the user account of ", account.Name)
		gogit.UpdateLocalCopies(repoNames, config, &account)
	}

	BulkCron := cron.New()
	BulkCron.AddFunc(config.UpdateInterval, func() {
		for _, account := range config.Accounts {
			if account.BackupRepos {
				logr.Printf("[API] Found user: %v", account)
				repoNames = gitapi.GetRepoList(&account)
				logr.Info("[main] Found ", len(repoNames), " repositories on the account of ", account.Name)
				gogit.UpdateLocalCopies(repoNames, config, &account)
			} else {
				logr.Infof("[main] Skipping account %v because it's not set to backup repositories", account.Name)
				continue
			}
		}
	})
	BulkCron.Start()

	logr.Info("The initial run/backup cycle has completed")
	logr.Info("The cron jobs have been setup to run in the background ...")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	logr.Info("Received an interrupt signal, stopping the cron jobs ...")

	BulkCron.Stop()
	logr.Info("Cron jobs stopped, exiting ...")
	os.Exit(0)
}
