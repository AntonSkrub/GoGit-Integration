package goGit

import (
	"GoGit-Integration/pkg/config"
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func UpdateLocalCopies(names []string, config *config.Config) {
	url := "https://github.com/" + config.OrgaName

	for i := 0; i < len(names); i++ {
		fmt.Printf("Value of i currently is: %v\n", i)
		if i >= 3 {
			fmt.Println("we break here")
			break
		}

		// Open the repository at the give path
		r, err := git.PlainOpen(config.OutputPath + names[i])
		if err != nil {
			if err == git.ErrRepositoryNotExists {
				fmt.Printf("Error: Local copy of Repository %v not found, creating one now!\n", names[i])
				Clone(names[i], config)
			} else {
				fmt.Printf("Error: %v\n", err)
			}
			continue
		}

		// Retrieve the working directory for the repository
		w, err := r.Worktree()
		CheckIfError(err)

		// Pull the latest changes from the origin and merge into the current branch
		Info("git pull origin from %s to %s", url+names[i]+".git", config.OutputPath+names[i])
		err = w.Pull(&git.PullOptions{
			Auth: &http.BasicAuth{
				Username: config.OrgaName,
				Password: config.OrgaToken,
			},
			RemoteName: "origin",
			Progress:   os.Stdout,
		})
		if err != nil {
			if err == git.NoErrAlreadyUpToDate {
				fmt.Printf("Error: Repository already up to date\n")
			} else {
				fmt.Printf("Error: %v\n", err)
				return
			}
			continue
		}

		// ... retrieving the branch being pointed by HEAD
		ref, err := r.Head()
		CheckIfError(err)
		fmt.Println(ref)

		if !config.EnableLog {
			commit, err := r.CommitObject(ref.Hash())
			fmt.Printf("Commit: %+v\n", commit)
			CheckIfError(err)

			fmt.Println(commit)
		} else {
			GetLog(r, ref)
		}
	}
}

func Clone(name string, config *config.Config) {
	url := "https://github.com/" + config.OrgaName

	// Clone the given repository to the given directory
	Info("git clone %s to %s", url+name+".git", config.OutputPath+name)
	r, err := git.PlainClone(config.OutputPath+name, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: config.OrgaName,
			Password: config.OrgaToken,
		},
		URL:      url + name + ".git",
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	ref, err := r.Head()
	CheckIfError(err)
	fmt.Println(ref)

	if !config.EnableLog {
		commit, err := r.CommitObject(ref.Hash())
		fmt.Printf("Commit: %+v\n", commit)
		CheckIfError(err)

		fmt.Println(commit)
	} else {
		GetLog(r, ref)
	}

	fmt.Printf("finished cloning the %v Repository to %v\n", name, config.OutputPath+name)
}

func GetLog(r *git.Repository, ref *plumbing.Reference) {
	since := time.Now().AddDate(0, 0, -1)
	until := time.Now()
	cIter, err := r.Log(&git.LogOptions{
		From:  ref.Hash(),
		Since: &since,
		Until: &until,
	})
	CheckIfError(err)

	// ... just iterates over the commits, printing it
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)
		return nil
	})
	CheckIfError(err)
}
