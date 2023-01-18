package gitapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/AntonSkrub/GoGit-Integration/pkg/config"

	logr "github.com/sirupsen/logrus"
)

func GetList(config *config.Config) []string {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/orgs/"+config.OrgaName+"repos", nil)
	if err != nil {
		logr.Errorf("[API] failed creating the request: %v\n", err)
	}
	req.Header.Set("Authorization", "Bearer "+config.OrgaToken)

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
	var repos []map[string]interface{}
	err = json.Unmarshal(body, &repos)
	if err != nil {
		logr.Errorf("[API] failed unmarshalling the json: %v\n", err)
	}

	i := 0
	var repoNames []string
	for _, repo := range repos {
		i++
		name, ok := repo["name"].(string)
		if !ok {
			logr.Errorf("[API] failed converting the repository name to string: %v\n", err)
		} else {
			repoNames = append(repoNames, name)
		}
	}
	logr.Printf("[API] Found %v Repositories!", i)
	return repoNames
}
