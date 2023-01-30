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
	Owner    Owner  `json:"owner"`
}

type Owner struct {
	Login string `json:"login"`
	Type  string `json:"type"`
}

func GetRepoList(account *config.Account) []Repository {
	token, reqUrl := "", ""
	var err error

	reqUrl = buildURL("https://api.github.com/user/repos", "type", account.Option)
	token = account.Token

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
	if resp.StatusCode != http.StatusOK {
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
