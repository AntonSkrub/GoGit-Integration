package gitapi

import (
	"GoGit-Integration/pkg/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetList(config *config.Config) []string {

	req, err := http.NewRequest("GET", "https://api.github.com/orgs/"+config.OrgaName+"/repos", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	req.Header.Set("Authorization", "Bearer "+config.OrgaToken)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// unmarshal the json and get the name parameter for each repo
	var repos []map[string]interface{}
	err = json.Unmarshal(body, &repos)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	i := 0
	var repoNames []string
	fmt.Println("Repositories Names:")
	for _, repo := range repos {
		i++
		repoNames = append(repoNames, repo["name"].(string))
		fmt.Printf("%v. %v\n", i, repo["name"])
	}
	return repoNames
}
