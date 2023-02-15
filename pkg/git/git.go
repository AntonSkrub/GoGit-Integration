package git

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gogit-integration/pkg/config"
	"gogit-integration/pkg/gitapi"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"

	logr "github.com/sirupsen/logrus"
)

func UpdateLocalCopies(repos []gitapi.Repository, config *config.Config, account *config.Account) {
	for _, repo := range repos {
		if account.ValidateName && !strings.Contains(repo.FullName, account.Name) {
			logr.Infof("[Git] Skipping repository %v because it doesn't contain the account name %v", repo.FullName, account.Name)
			continue
		}

		r, err := git.PlainOpen(filepath.Join(config.OutputPath, repo.FullName))
		if err != nil {
			if err == git.ErrRepositoryNotExists {
				logr.Errorf("[Git] Couldn't find a local copy of Repository %v", repo.FullName)
				Clone(repo.FullName, config, account)
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
		auth := buildAuth(account)

		logr.Infof("[Git] Pulling the latest changes from the origin of %v", repo.FullName)
		err = w.Pull(&git.PullOptions{
			RemoteName:   "origin",
			SingleBranch: false,
			Auth:         auth,
			Progress:     os.Stdout,
		})
		if err != nil {
			if err == git.NoErrAlreadyUpToDate {
				logr.Errorf("[Git] Repository %v already up to date", repo.FullName)
			} else {
				logr.Errorf("[Git] failed pulling the repository: %v\n", err)
				return
			}
			continue
		}

		if config.ListReferences {
			AccessRepo(r, config, account)
		}
		logr.Infof("[Git] Finished updating the %v repository", repo.FullName)
	}
}

func Clone(name string, config *config.Config, account *config.Account) {
	url, err := url.JoinPath("https://github.com", name+".git")
	if err != nil {
		logr.Errorf("[GoGit] failed creating the url: %v\n", err)
		return
	}

	auth := buildAuth(account)
	// Clone the repository to the given directory
	path := filepath.Join(config.OutputPath, name)
	logr.Infof("[GoGit] Cloning the %v repository to %v", url, path)
	r, err := git.PlainClone(path, false, &git.CloneOptions{
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

	if config.ListReferences {
		AccessRepo(r, config, account)
	}
	logr.Infof("[GoGit] finished cloning the %v repository to %v", name, path)
}

func buildAuth(account *config.Account) *http.BasicAuth {
	auth := &http.BasicAuth{
		Username: account.Name,
		Password: account.Token,
	}
	return auth
}
