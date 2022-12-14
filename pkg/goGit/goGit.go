package goGit

import (
	"GoGit-Integration/pkg/config"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"

	logr "github.com/sirupsen/logrus"
)

func UpdateLocalCopies(names []string, config *config.Config) {
	for i := 0; i < len(names); i++ {
		fmt.Printf("Value of i currently is: %v\n", i)
		if i >= 5 {
			fmt.Println("we break here")
			break
		}

		// Open the repository at the give path
		r, err := git.PlainOpen(config.OutputPath + names[i])
		if err != nil {
			if err == git.ErrRepositoryNotExists {
				logr.Errorf("[GoGit] Couldn't find a local copy of Repository %v", names[i])
				Clone(names[i], config)
			} else {
				logr.Errorf("[GoGit] failed opening the repository: %v\n", err)
			}
			continue
		}
		// Retrieve the working directory for the repository
		w, err := r.Worktree()
		if err != nil {
			logr.Errorf("[GoGit] failed getting the working directory: %v\n", err)
			return
		}

		// Pull the latest changes from the origin and merge into the current branch
		logr.Infof("[GoGit] Pulling the latest changes from the origin of %v", names[i])
		err = w.Pull(&git.PullOptions{
			RemoteName:   "origin",
			SingleBranch: false,
			Auth: &http.BasicAuth{
				Username: config.OrgaName,
				Password: config.OrgaToken,
			},
			Progress: os.Stdout,
		})
		if err != nil {
			if err == git.NoErrAlreadyUpToDate {
				logr.Errorf("[GoGit] Repository %v already up to date", names[i])
			} else {
				logr.Errorf("[GoGit] failed pulling the repository: %v\n", err)
				return
			}
			continue
		}

		if config.ListRefereces {
			ListRefs(r, config)
		}

		if config.LogLatestCommits {
			GetLog(r, config)
		}
		logr.Infof("[GoGit] Finished updating the %v repository", names[i])
	}
}

func Clone(name string, config *config.Config) {
	url := "https://github.com/" + config.OrgaName

	// Clone the repository to the given directory
	logr.Infof("[GoGit] Cloning the %v repository to %v", url+name+".git", config.OutputPath+name)
	r, err := git.PlainClone(config.OutputPath+name, false, &git.CloneOptions{
		URL:          url + name + ".git",
		RemoteName:   "origin",
		SingleBranch: false,
		NoCheckout:   false,
		Auth: &http.BasicAuth{
			Username: config.OrgaName,
			Password: config.OrgaToken,
		},
		Progress: os.Stdout,
	})
	if err != nil {
		logr.Errorf("[GoGit] failed cloning the repository: %v\n", err)
		return
	}

	if config.ListRefereces {
		ListRefs(r, config)
	}

	if config.LogLatestCommits {
		GetLog(r, config)
	}
	logr.Infof("[GoGit] finished cloning the %v repository to %v", name, config.OutputPath+name)
}
