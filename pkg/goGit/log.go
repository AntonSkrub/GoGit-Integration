package gogit

import (
	"strings"

	"github.com/AntonSkrub/GoGit-Integration/pkg/config"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	logr "github.com/sirupsen/logrus"
)

func AccessRepo(r *git.Repository, config *config.Config, account *config.Account) {
	remote, err := r.Remote("origin")
	if err != nil {
		logr.Errorf("[GoGit] failed getting the remote: %v\n", err)
		return
	}
	refList, err := remote.List(&git.ListOptions{
		Auth: &http.BasicAuth{
			Username: account.Name,
			Password: account.Token,
		},
	})
	if err != nil {
		logr.Errorf("[GoGit] failed listing the remote: %v\n", err)
		return
	}

	if config.ListReferences {
		ListRefs(refList)
	}
	if config.LogCommits {
		GetLog(r, refList)
	}
}

func ListRefs(refList []*plumbing.Reference) {
	branchRefPrefix := "refs/heads/"
	tagRefPrefix := "refs/tags/"
	branchList := make([]string, 0)
	tagList := make([]string, 0)

	for _, ref := range refList {
		refName := ref.Name().String()
		if !strings.HasPrefix(refName, branchRefPrefix) {
			if strings.HasPrefix(refName, tagRefPrefix) {
				tagList = append(tagList, refName)
			}
			continue
		}
		branchList = append(branchList, refName)
	}

	for _, branch := range branchList {
		branchName := branch[len(branchRefPrefix):]
		logr.Infof("[Git] found branch: %s", branchName)
	}

	if len(tagList) == 0 {
		logr.Info("[Git] no tags found")
		return
	}

	for _, tag := range tagList {
		tagName := tag[len(tagRefPrefix):]
		logr.Infof("[Git] found tag: %s", tagName)
	}
}

func GetLog(r *git.Repository, refList []*plumbing.Reference) {
	branchRefPrefix := "refs/heads/"
	for _, ref := range refList {
		refName := ref.Name().String()
		if !strings.HasPrefix(refName, branchRefPrefix) {
			continue
		}
		branchName := refName[len(branchRefPrefix):]

		commit, err := r.CommitObject(ref.Hash())
		if err != nil {
			logr.Errorf("[Git] failed getting the commit object: %v\n", err)
			return
		}
		// log the branch name
		logr.Infof("[Git] Get latest commit on %s branch", branchName)
		// log the commit
		logr.Infof("[Git] commit: %s", commit.Hash)
		logr.Infof("[Git] author: %s", commit.Author)
		logr.Infof("[Git] message: %s", commit.Message)
	}
}
