package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/AntonSkrub/GoGit-Integration/pkg/config"
	"github.com/AntonSkrub/GoGit-Integration/pkg/gitapi"
	"github.com/AntonSkrub/GoGit-Integration/pkg/gogit"
	logr "github.com/sirupsen/logrus"
)

func main() {
	logr.Infoln("Grabbing configuration...")
	config := config.GetConfig()
	logr.SetLevel(logr.Level(config.LogLevel))

	for _, orga := range config.Organizations {
		logr.Printf("[API] Found organization: %v", orga.Name)
		orgaRepoNames := gitapi.GetRepoList(&orga, nil)
		logr.Info("[main] Found ", len(orgaRepoNames), " repositories in the organization ", orga.Name)
		gogit.UpdateLocalCopies(orgaRepoNames, config, &orga, nil)
	}

	// orgaRepoNames := gitapi.GetRepoList(config, nil)
	// logr.Info("[main] Found ", len(orgaRepoNames), " repositories in the organization")

	// gogit.UpdateLocalCopies(orgaRepoNames, config, nil)

	// loop through the users in the config and log each users name to the console
	for _, user := range config.Users {
		logr.Printf("[API] Found user: %v", user)
		userRepoNames := gitapi.GetRepoList(nil, &user)
		logr.Info("[main] Found ", len(userRepoNames), " repositories on the user account of ", user.Name)
		gogit.UpdateLocalCopies(userRepoNames, config, nil, &user)
	}

	// UpdateInterval := cron.New()
	// UpdateInterval.AddFunc(config.UpdateInterval, func() {
	// 	gogit.UpdateLocalCopies(orgaRepoNames, config, nil)
	// })
	// go UpdateInterval.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	logr.Info("Received an interrupt signal, stopping the cron jobs ...")

	// UpdateInterval.Stop()
	logr.Info("Cron jobs stopped, exiting ...")
	os.Exit(0)
}
