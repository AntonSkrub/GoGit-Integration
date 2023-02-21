package main

import (
	"flag"
	"fmt"
	"gogit-integration/pkg/config"
	"gogit-integration/pkg/git"
	"gogit-integration/pkg/gitapi"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/robfig/cron/v3"
	logr "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const version = "1.0.0"

const logname = "github-backup.log"
const logFilePath = "./logs/"

func main() {
	showHelp := flag.Bool("help", false, "Show general information about the GoGit-Integration.")
	showConfigHelp := flag.Bool("confighelp", false, "Displays all configuration options.")
	setConfigPath := flag.String("config", "./config", "Sets the path to the configuration file.")
	debugMode := flag.Bool("debug", false, "Enables the debug loglevel for this lifecycle.")
	flag.Parse()

	if *showHelp {
		printHelp()
	}
	if *showConfigHelp {
		printConfigExplanation()
	}
	if *setConfigPath != "" {
		config.SetConfigPath(*setConfigPath)
	}

	if *showHelp || *showConfigHelp {
		return
	}

	logr.Infoln("Grabbing configuration...")
	config := config.GetConfig()
	if *debugMode {
		config.LogLevel = uint32(logr.DebugLevel)
	}

	// set a lumberjack configuration
	logr.Debugln("Configuring lumberjack logger ...")
	logPath := filepath.Join(logFilePath, logname)
	logger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    int(config.MaxLogLength), // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
		Compress:   false,
	}

	mw := io.MultiWriter(logger, os.Stdout)
	logr.SetLevel(logr.Level(config.LogLevel))
	logr.SetOutput(mw)
	logr.Debugln("The application is running in debug mode")

	repoNames := []gitapi.Repository{}
	for _, account := range config.Accounts {
		if !account.BackupRepos {
			logr.Infof("[main] Skipping the account %v because not configured to be backed up", account.Name)
			continue
		}

		logr.Printf("[main] Found account: %v", account.Name)
		repoNames = gitapi.GetRepoList(&account)
		if len(repoNames) == 0 {
			logr.Infof("[main] Couldn't find any repositories on the account %v", account.Name)
			continue
		}
		logr.Infof("[main] Found %v repositories on the user account of %v", len(repoNames), account.Name)
		git.UpdateLocalCopies(repoNames, config, &account)
	}

	BulkCron := cron.New()
	BulkCron.AddFunc(config.UpdateInterval, func() {
		for _, account := range config.Accounts {
			if account.BackupRepos {
				logr.Printf("[API] Found user: %v", account)
				repoNames = gitapi.GetRepoList(&account)
				if len(repoNames) == 0 {
					logr.Infof("[main] No repositories found on the account of %v", account.Name)
					continue
				}
				logr.Info("[main] Found ", len(repoNames), " repositories on the account of ", account.Name)
				git.UpdateLocalCopies(repoNames, config, &account)
			} else {
				logr.Infof("[main] Skipping account %v because it's not set to backup repositories", account.Name)
				continue
			}
		}
	})
	BulkCron.Start()

	logr.Info("Completed the initial backup cycle")
	logr.Info("Cron jobs set up to run in the background ...")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	logr.Info("Received an interrupt signal, stopping the cron jobs ...")

	BulkCron.Stop()
	logr.Info("Cron jobs stopped, exiting ...")
	os.Exit(0)
}

func printHelp() {
	fmt.Println("=====================================")
	fmt.Printf("GoGit-Integration %v - developed by Anton Paul\n", version)
	fmt.Println("Golang based tool for backing up repositories a or multiple GitHub accounts.")
	fmt.Println("=====================================")
	fmt.Printf("Allowed starting flags:\n\n")
	fmt.Printf("-confighelp - Displays all configuration options.\n")
	fmt.Printf("-config - Sets the path to the configuration file.\n")
}

func printConfigExplanation() {
	fmt.Println("=====================================")
	fmt.Println("The configuration file must be named config.yml and placed in a config directory where the executable file is located.")
	fmt.Printf("The following parameters must be configured in a config file: \n \n")
	fmt.Printf("Accounts: A list of GitHub accounts to backup repositories from. \n")
	fmt.Printf("  Name: The name of the account. \n")
	fmt.Printf("  Token: The access-token for the account. \n")
	fmt.Printf("  Option: Added to the requested URL as a parameter. Defines which repositories to process. \n\t  Possible values are: all, owner, public, private, member. \n")
	fmt.Printf("  BackupRepos: A boolean value indicating whether the repositories of this account should be backed up. \n")
	fmt.Printf("  ValidateName: A boolean value indicating whether the repositories `full_name` has to contain the `Account.Name`. \n")
	fmt.Printf("OutputPath: The path where the repositories should be stored. \n")
	fmt.Printf("UpdateInterval: The interval in which the repositories should be updated. \n")
	fmt.Printf("  Is set in the cron syntax improved by spring, find more information \n  here https://spring.io/blog/2020/11/10/new-in-spring-5-3-improved-cron-expressions \n")
	fmt.Printf("ListReferences: A boolean indicating whether the references (branches, tags) should be listed. \n")
	fmt.Printf("LogLevel: An integrer representing the desired loglevel. \n")
	fmt.Printf("MaxLogLength: The maximum size of the log file in MB\n")
	fmt.Println("=====================================")
}
