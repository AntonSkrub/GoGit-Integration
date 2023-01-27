package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/AntonSkrub/GoGit-Integration/pkg/config"
	"github.com/AntonSkrub/GoGit-Integration/pkg/gitapi"
	"github.com/AntonSkrub/GoGit-Integration/pkg/gogit"
	"github.com/robfig/cron"
	logr "github.com/sirupsen/logrus"
)

func main() {
	logr.Infoln("Grabbing configuration...")
	config := config.GetConfig()
	logr.SetLevel(logr.Level(config.LogLevel))

	// loop through the organizations in the config and log each organizations name to the console
	orgaRepoNames := []gitapi.Repository{}
	for _, orga := range config.Organizations {
		if orga.BackupRepos {
			logr.Printf("[API] Found organization: %v", orga.Name)
			orgaRepoNames = gitapi.GetRepoList(&orga, nil)
			logr.Info("[main] Found ", len(orgaRepoNames), " repositories in the organization ", orga.Name)
			gogit.UpdateLocalCopies(orgaRepoNames, config, &orga, nil)
		} else {
			logr.Infof("[main] Skipping organization %v because it's not set to backup repositories", orga.Name)
			continue
		}
	}

	// loop through the users in the config and log each users name to the console
	userRepoNames := []gitapi.Repository{}
	for _, user := range config.Users {
		if user.BackupRepos {
			logr.Printf("[API] Found user: %v", user)
			userRepoNames = gitapi.GetRepoList(nil, &user)
			logr.Info("[main] Found ", len(userRepoNames), " repositories on the user account of ", user.Name)
			// gogit.UpdateLocalCopies(userRepoNames, config, nil, &user)
		} else {
			logr.Infof("[main] Skipping user %v because it's not set to backup repositories", user.Name)
			continue
		}
	}

	OrgaCron := cron.New()
	OrgaCron.AddFunc(config.UpdateInterval, func() {
		for _, orga := range config.Organizations {
			if orga.BackupRepos {
				logr.Printf("[API] Found organization: %v", orga.Name)
				orgaRepoNames = gitapi.GetRepoList(&orga, nil)
				logr.Info("[main] Found ", len(orgaRepoNames), " repositories in the organization ", orga.Name)
				gogit.UpdateLocalCopies(orgaRepoNames, config, &orga, nil)
			} else {
				logr.Infof("[main] Skipping organization %v because it's not set to backup repositories", orga.Name)
				continue
			}
		}
	})
	OrgaCron.Start()

	UserCron := cron.New()
	UserCron.AddFunc(config.UpdateInterval, func() {
		for _, user := range config.Users {
			if user.BackupRepos {
				logr.Printf("[API] Found user: %v", user)
				userRepoNames = gitapi.GetRepoList(nil, &user)
				logr.Info("[main] Found ", len(userRepoNames), " repositories on the user account of ", user.Name)
				gogit.UpdateLocalCopies(userRepoNames, config, nil, &user)
			} else {
				logr.Infof("[main] Skipping user %v because it's not set to backup repositories", user.Name)
				continue
			}
		}
	})
	UserCron.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	logr.Info("Received an interrupt signal, stopping the cron jobs ...")

	UserCron.Stop()
	OrgaCron.Stop()
	logr.Info("Cron jobs stopped, exiting ...")
	os.Exit(0)
}
