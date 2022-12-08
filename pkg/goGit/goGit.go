package goGit

import (
	"GoGit-Integration/pkg/config"
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func UpdateLocalCopies(names []string, config *config.Config) {
	url := "https://github.com/" + config.OrgaName + "/"
	dir := "../Go-Test/"

	for i := 0; i < len(names); i++ {
		fmt.Printf("Value of i currently is: %v\n", i)
		if i >= 3 {
			fmt.Println("we break here")
			break
		}

		// Open the repository at the give path
		r, err := git.PlainOpen(dir + names[i])
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
		Info("git pull origin from %s to %s", url+names[i]+".git", dir+names[i])
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

		//---
		// ... retrieving the branch being pointed by HEAD
		ref, err := r.Head()
		CheckIfError(err)
		fmt.Println(ref)

		// since := time.Date(2022, 10, 1, 0, 0, 0, 0, time.UTC)
		// until := time.Date(2022, 12, 9, 0, 0, 0, 0, time.UTC)
		since := time.Now().AddDate(0, 0, -2)
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
		//---
	}
}

func Clone(name string, config *config.Config) {
	url := "https://github.com/" + config.OrgaName + "/"
	dir := "../Go-Test/"

	// Clone the given repository to the given directory
	Info("git clone %s to %s", url+name+".git", dir+name)
	r, err := git.PlainClone(dir+name, false, &git.CloneOptions{
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

	//---
	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)
	fmt.Println(ref)

	// since := time.Date(2022, 10, 1, 0, 0, 0, 0, time.UTC)
	// until := time.Date(2022, 12, 9, 0, 0, 0, 0, time.UTC)
	since := time.Now().AddDate(0, 0, -2)
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
	//---

	fmt.Printf("finished cloning the %v Repository to %v\n", name, dir+name)
}
