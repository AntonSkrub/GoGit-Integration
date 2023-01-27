package gitapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/AntonSkrub/GoGit-Integration/pkg/config"

	logr "github.com/sirupsen/logrus"
)

type Repository struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    string `json:"owner.login"`
}

func GetRepoList(orga *config.Account, user *config.Account) []Repository {
	token, reqUrl := "", ""
	var err error
	if user != nil {
		reqUrl = buildURL("https://api.github.com/user/repos", "affiliation", user.Option)
		token = user.Token
	} else {
		baseUrl, err := url.JoinPath("https://api.github.com/orgs/", orga.Name, "repos")
		if err != nil {
			logr.Errorf("[API] failed creating the url: %v\n", err)
		}
		reqUrl = buildURL(baseUrl, "type", orga.Option)
		token = orga.Token
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
	var repos []Repository
	err = json.Unmarshal(body, &repos)
	if err != nil {
		logr.Errorf("[API] failed unmarshalling the json: %v\n", err)
	}

	logr.Printf("[API] Found %v Repositories!", len(repos))
	return repos
}

func buildURL(baseURL string, paramType string, param string) string {
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
