package gitapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/AntonSkrub/GoGit-Integration/pkg/config"

	logr "github.com/sirupsen/logrus"
)

type Repo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    string `json:"owner.login"`
}

func GetRepoList(config *config.Config, user *config.User) []string {
	token, reqUrl := "", ""
	var err error
	if user != nil {
		reqUrl = buildURL("https://api.github.com/user/repos", "affiliation", user.Affiliation)
		token = user.Token
	} else {
		baseUrl, err := url.JoinPath("https://api.github.com/orgs/", config.OrgaName, "repos")
		if err != nil {
			logr.Errorf("[API] failed creating the url: %v\n", err)
		}
		reqUrl = buildURL(baseUrl, "type", config.OrgaRepoType)
		token = config.OrgaToken
	}

	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		logr.Errorf("[API] failed creating the request: %v\n", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		logr.Errorf("[API] failed sending the request: %v\n", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logr.Errorf("[API] failed reading the response body: %v\n", err)
	}

	// Unmarshal the json response to get the repository names
	var repos []Repo
	err = json.Unmarshal(body, &repos)
	if err != nil {
		logr.Errorf("[API] failed unmarshalling the json: %v\n", err)
	}

	i := 0
	var repoNames []string
	for _, repo := range repos {
		i++
		repoNames = append(repoNames, repo.FullName)
	}
	logr.Printf("[API] Found %v Repositories!", i)
	return repoNames
}

func buildURL(baseURL string, paramType string, param string) string {
	fmt.Printf("using paramType: %v with param: %v", paramType, param)
	url, err := url.Parse(baseURL)
	if err != nil {
		logr.Errorf("[API] failed creating the url: %v\n", err)
	}
	q := url.Query()
	q.Add(paramType, param)
	url.RawQuery = q.Encode()
	urlString := url.String()
	return urlString
}
