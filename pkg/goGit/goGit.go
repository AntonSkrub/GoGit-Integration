package gogit

import (
	"net/url"
	"os"
	"path/filepath"

	"github.com/AntonSkrub/GoGit-Integration/pkg/config"
	"github.com/AntonSkrub/GoGit-Integration/pkg/gitapi"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"

	logr "github.com/sirupsen/logrus"
)

func UpdateLocalCopies(repos []gitapi.Repository, config *config.Config, orga *config.Organization, user *config.User) {
	for _, repo := range repos {
		r, err := git.PlainOpen(filepath.Join(config.OutputPath, repo.FullName))
		if err != nil {
			if err == git.ErrRepositoryNotExists {
				logr.Errorf("[Git] Couldn't find a local copy of Repository %v", repo.FullName)
				Clone(repo.FullName, config, orga, user)
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
		auth := buildAuth(orga, user)

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

		if config.ListReferences || config.LogCommits {
			AccessRepo(r, config)
		}
		logr.Infof("[Git] Finished updating the %v repository", repo.FullName)
	}
}

func Clone(name string, config *config.Config, orga *config.Organization, user *config.User) {
	url, err := url.JoinPath("https://github.com", name+".git")
	if err != nil {
		logr.Errorf("[GoGit] failed creating the url: %v\n", err)
		return
	}

	auth := buildAuth(orga, user)
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

	if config.ListReferences || config.LogCommits {
		AccessRepo(r, config)
	}
	logr.Infof("[GoGit] finished cloning the %v repository to %v", name, path)
}

func buildAuth(orga *config.Organization, user *config.User) *http.BasicAuth {
	var auth *http.BasicAuth
	auth = &http.BasicAuth{
		Username: orga.Name,
		Password: orga.Token,
	}

	if user != nil {
		auth = &http.BasicAuth{
			Username: user.Name,
			Password: user.Token,
		}
	}
	return auth
}
