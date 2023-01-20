package gogit

import (
	"os"
	"path/filepath"

	"github.com/AntonSkrub/GoGit-Integration/pkg/config"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"

	logr "github.com/sirupsen/logrus"
)

func UpdateLocalCopies(names []string, config *config.Config, user *config.User) {
	for i := 0; i < len(names); i++ {
		if i >= 10 {
			logr.Info("Maximum number of repositories reached, exiting...")
			break
		}
		// Open the repository at the give path

		r, err := git.PlainOpen(filepath.Join(config.OutputPath, names[i]))
		if err != nil {
			if err == git.ErrRepositoryNotExists {
				logr.Errorf("[Git] Couldn't find a local copy of Repository %v", names[i])

				if user != nil {
					Clone(names[i], config, user)
				} else {
					Clone(names[i], config, nil)
				}
			} else {
				logr.Errorf("[Git] failed opening the repository: %v\n", err)
			}
			continue
		}
		// Retrieve the working directory for the repository
		w, err := r.Worktree()
		if err != nil {
			logr.Errorf("[Git] failed getting the working directory: %v\n", err)
			return
		}

		// Pull the latest changes from the origin and merge into the current branch
		var auth *http.BasicAuth
		if user != nil {
			auth = &http.BasicAuth{
				Username: user.Name,
				Password: user.Token,
			}
		} else {
			auth = &http.BasicAuth{
				Username: config.OrgaName,
				Password: config.OrgaToken,
			}
		}

		logr.Infof("[Git] Pulling the latest changes from the origin of %v", names[i])
		err = w.Pull(&git.PullOptions{
			RemoteName:   "origin",
			SingleBranch: false,
			Auth:         auth,
			Progress:     os.Stdout,
		})
		if err != nil {
			if err == git.NoErrAlreadyUpToDate {
				logr.Errorf("[Git] Repository %v already up to date", names[i])
			} else {
				logr.Errorf("[Git] failed pulling the repository: %v\n", err)
				return
			}
			continue
		}

		if config.ListReferences || config.LogCommits {
			AccessRepo(r, config)
		}
		logr.Infof("[Git] Finished updating the %v repository", names[i])
	}
}

func Clone(name string, config *config.Config, user *config.User) {
	url := "https://github.com/" + name + ".git"

	var auth *http.BasicAuth
	if user != nil {
		auth = &http.BasicAuth{
			Username: user.Name,
			Password: user.Token,
		}
	} else {
		auth = &http.BasicAuth{
			Username: config.OrgaName,
			Password: config.OrgaToken,
		}
	}

	// Clone the repository to the given directory
	logr.Infof("[GoGit] Cloning the %v repository to %v", url, config.OutputPath+name)
	r, err := git.PlainClone(config.OutputPath+name, false, &git.CloneOptions{
		URL:          url,
		RemoteName:   "origin",
		SingleBranch: false,
		NoCheckout:   false,
		Auth:         auth,
		Progress:     os.Stdout,
	})
	if err != nil {
		logr.Errorf("[GoGit] failed cloning the repository: %v\n", err)
		return
	}

	if config.ListReferences || config.LogCommits {
		AccessRepo(r, config)
	}
	logr.Infof("[GoGit] finished cloning the %v repository to %v", name, config.OutputPath+name)
}
