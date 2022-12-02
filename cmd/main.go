package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func main() {
	fmt.Println("Erstmal sachte moin")
	token := "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"
	url := "https://github.com/Avanis-GmbH/"
	dir := "../Go-Test/"

	// create a string array with 4 values
	var arr [4]string
	arr[0] = "Flip-Catalog"
	arr[1] = "Form-Transmitter"
	arr[2] = "Hugo-Dani-Apart"
	arr[3] = "Go-Dust-Vacuum"
	// iterage over the array
	for i := 0; i < len(arr); i++ {
		Info("git clone %s %s", url+arr[i]+".git", dir+arr[i])
		// Clones the repository into the given dir, just as a normal git clone does
		r, err := git.PlainClone(dir+arr[i], false, &git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: "ce moi",
				Password: token,
			},
			URL:      url + arr[i] + ".git",
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
		fmt.Printf("finished cloning for the %v. repository", i)
	}

	fmt.Println("Ente gut alles gut")
}
