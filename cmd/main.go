package main

import (
	"GoGit-Integration/pkg/config"
	"GoGit-Integration/pkg/gitapi"
	"GoGit-Integration/pkg/goGit"
	"fmt"
	"os"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	names := gitapi.GetList(config)

	fmt.Println("Repositories Paths:")
	for i := 0; i < len(names); i++ {
		fmt.Printf("%v. %v/%v\n", i, config.OrgaName, names[i])
	}

	goGit.Clone(names, config)
}
