package gitapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetList() []string {
	req, err := http.NewRequest("GET", "https://api.github.com/orgs/Avanis-GmbH/repos", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	req.Header.Set("Authorization", "Bearer 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")

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
