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

func GetRepoList(account *config.Account) []Repository {
	token, reqUrl := "", ""
	var err error
	// has error on GoSUCK with current config
	if account.Type == "organization" {
		baseUrl, err := url.JoinPath("https://api.github.com/orgs/", account.Name, "repos")
		if err != nil {
			logr.Errorf("[API] failed creating the url: %v\n", err)
		}
		reqUrl = buildURL(baseUrl, "type", account.Option)
		token = account.Token
	} else if account.Type == "user" {
		reqUrl = buildURL("https://api.github.com/user/repos", "affiliation", account.Option)
		token = account.Token
	}
	
	// works with current config
	// reqUrl = "https://api.github.com/user/repos"
	// token = account.Token


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
	if resp.StatusCode != 200 {
		logr.Errorf("[API] Error %d: %s", resp.StatusCode, resp.Status)
		return nil
	}

	// Unmarshal the json response to get the repository names
	var repos []Repository
	err = json.Unmarshal(body, &repos)
	if err != nil {
		logr.Errorf("[API] failed unmarshalling the json: %v\n", err)
	}

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
