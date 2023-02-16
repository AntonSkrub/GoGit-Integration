package git

import (
	"gogit-integration/pkg/config"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	logr "github.com/sirupsen/logrus"
)

func AccessRepo(r *git.Repository, config *config.Config, account *config.Account) {
	remote, err := r.Remote("origin")
	if err != nil {
		logr.Errorf("[Git] failed getting the remote: %v\n", err)
		return
	}
	refList, err := remote.List(&git.ListOptions{
		Auth: &http.BasicAuth{
			Username: account.Name,
			Password: account.Token,
		},
	})
	if err != nil {
		logr.Errorf("[Git] failed listing the remote: %v\n", err)
		return
	}

	if config.ListReferences {
		ListRefs(refList)
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
