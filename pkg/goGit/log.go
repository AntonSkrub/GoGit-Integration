package goGit

import (
	"GoGit-Integration/pkg/config"
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	logr "github.com/sirupsen/logrus"
)

func ListRefs(r *git.Repository, config *config.Config) {
	fmt.Println("------ remote branch references --------")
	remote, err := r.Remote("origin")
	if err != nil {
		logr.Errorf("[GoGit] failed getting the remote: %v\n", err)
		return
	}
	refList, err := remote.List(&git.ListOptions{
		Auth: &http.BasicAuth{
			Username: config.OrgaName,
			Password: config.OrgaToken,
		},
	})
	if err != nil {
		logr.Errorf("[GoGit] failed listing the remote: %v\n", err)
		return
	}
	branchRefPrefix := "refs/heads/"
	for _, ref := range refList {
		refName := ref.Name().String()
		if !strings.HasPrefix(refName, branchRefPrefix) {
			continue
		}
		branchName := refName[len(branchRefPrefix):]
		fmt.Println(branchName)
	}

	fmt.Println("------ tag references --------")
	tagList := make([]string, 0)
	tagRefPrefix := "refs/tags/"
	for _, ref := range refList {
		refName := ref.Name().String()
		if !strings.HasPrefix(refName, tagRefPrefix) {
			continue
		}
		tagList = append(tagList, refName)
	}
	if len(tagList) == 0 {
		fmt.Println("--- No tags found ---")
	} else {
		for _, tag := range tagList {
			tagName := tag[len(tagRefPrefix):]
			fmt.Println(tagName)
		}
	}
}

func GetLog(r *git.Repository, config *config.Config) {
	remote, err := r.Remote("origin")
	if err != nil {
		logr.Errorf("[GoGit] failed getting the remote: %v\n", err)
		return
	}
	refList, err := remote.List(&git.ListOptions{
		Auth: &http.BasicAuth{
			Username: config.OrgaName,
			Password: config.OrgaToken,
		},
	})
	if err != nil {
		logr.Errorf("[GoGit] failed listing the remote: %v\n", err)
		return
	}
	branchRefPrefix := "refs/heads/"
	for _, ref := range refList {
		refName := ref.Name().String()
		if !strings.HasPrefix(refName, branchRefPrefix) {
			continue
		}
		branchName := refName[len(branchRefPrefix):]

		commit, err := r.CommitObject(ref.Hash())
		if err != nil {
			logr.Errorf("[GoGit] failed getting the commit object: %v\n", err)
			return
		}
		logr.Infof("[GoGit] Get latest commit %v on branch\n", branchName)
		fmt.Printf("Git checkout: %v\n", commit)
	}
}
