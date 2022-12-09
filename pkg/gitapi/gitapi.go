package gitapi

import (
	"GoGit-Integration/pkg/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	logr "github.com/sirupsen/logrus"
)

func GetList(config *config.Config) []string {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/orgs/"+config.OrgaName+"repos", nil)
	if err != nil {
		logr.Errorf("[GitAPI] failed creating the request: %v\n", err)
	}
	req.Header.Set("Authorization", "Bearer "+config.OrgaToken)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		logr.Errorf("[GitAPI] failed sending the request: %v\n", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logr.Errorf("[GitAPI] failed reading the response body: %v\n", err)
	}

	// unmarshal the json and get the name parameter for each repo
	var repos []map[string]interface{}
	err = json.Unmarshal(body, &repos)
	if err != nil {
		logr.Errorf("[GitAPI] failed unmarshalling the json: %v\n", err)
	}

	i := 0
	var repoNames []string
	logr.Println("[GitAPI] Found Repositories:")
	for _, repo := range repos {
		i++
		name, ok := repo["name"].(string)
		if !ok {
			logr.Errorf("[GitAPI] failed converting the repository name to string: %v\n", err)
		} else {
			repoNames = append(repoNames, name)
			fmt.Printf("%v. %v%v\n", i, config.OrgaName, repo["name"])
		}
	}
	return repoNames
}
