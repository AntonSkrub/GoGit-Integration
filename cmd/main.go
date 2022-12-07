package main

import (
	"GoGit-Integration/pkg/config"
	"GoGit-Integration/pkg/gitapi"
	"GoGit-Integration/pkg/goGit"
	"fmt"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	names := gitapi.GetList(config)

	fmt.Println("Repositories Paths:")
	for i := 0; i < len(names); i++ {
		fmt.Printf("%v. Avanis-GmbH/%v\n", i, names[i])
	}

	goGit.Clone(names, config)
}
