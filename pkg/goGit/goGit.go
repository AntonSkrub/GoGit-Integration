package goGit

import (
	"GoGit-Integration/pkg/config"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func Clone(names []string, config *config.Config) {
	url := "https://github.com/Avanis-GmbH/"
	dir := "../Go-Test/"

	for i := 0; i < len(names); i++ {

		Info("git clone %s %s", url+names[i]+".git", dir+names[i])
		// Clones the repository into the given dir, just as a normal git clone does
		r, err := git.PlainClone(dir+names[i], false, &git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: "Avanis-GmbH",
				Password: config.OrgaToken,
			},
			URL:      url + names[i] + ".git",
			Progress: os.Stdout,
		})
		CheckIfError(err)

		// ... retrieving the branch being pointed by HEAD
		ref, err := r.Head()
		CheckIfError(err)
		// ... retrieving the commit object
		commit, err := r.CommitObject(ref.Hash())
		CheckIfError(err)

		fmt.Println(commit)
		fmt.Printf("finished cloning for the %v. repository\n", i+1)
	}
}
