package goGit

import (
	"GoGit-Integration/pkg/config"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func Clone(names []string, config *config.Config) {
	url := "https://github.com/" + config.OrgaName + "/"
	dir := "../Go-Test/"

	for i := 0; i < len(names); i++ {
		fmt.Printf("I is :%v\n", i)
		if i >= 10 {
			fmt.Println("we break here")
			break
		}
		log.Printf("git clone %s to %s", url+names[i]+".git", dir+names[i])
		// Clones the repository into the given dir, just as a normal git clone does
		r, err := git.PlainClone(dir+names[i], false, &git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: config.OrgaName,
				Password: config.OrgaToken,
			},
			URL:      url + names[i] + ".git",
			Progress: os.Stdout,
		})
		// if error is not nil, skip the repository and continue with the next one
		if err != nil {
			if err == git.ErrRepositoryAlreadyExists {
				fmt.Printf("Error: Repository already exists, updating it now\n")
				Update(names[i], config)
			} else {
				fmt.Printf("Error: %v\n", err)
			}
			continue
		}

		// ... retrieving the branch being pointed by HEAD
		ref, err := r.Head()
		CheckIfError(err)
		fmt.Println(ref)

		// since := time.Date(2022, 10, 1, 0, 0, 0, 0, time.UTC)
		// until := time.Date(2022, 12, 9, 0, 0, 0, 0, time.UTC)
		since := time.Now().AddDate(0, -2, 0)
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

		fmt.Printf("finished cloning for the %v. repository\n", i+1)
	}
}

func Update(name string, config *config.Config) {
	url := "https://github.com/" + config.OrgaName + "/"
	dir := "../Go-Test/"

	r, err := git.PlainOpen(dir + name)
	CheckIfError(err)

	w, err := r.Worktree()
	CheckIfError(err)

	Info("git pull origin from %s to %s", url+name+".git", dir+name)
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
			return
		} else {
			fmt.Printf("Error: %v\n", err)
			return
		}
	}

	ref, err := r.Head()
	CheckIfError(err)
	fmt.Println(ref)

	// since := time.Date(2022, 10, 1, 0, 0, 0, 0, time.UTC)
	// until := time.Date(2022, 12, 9, 0, 0, 0, 0, time.UTC)
	since := time.Now().AddDate(0, 0, -7)
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
