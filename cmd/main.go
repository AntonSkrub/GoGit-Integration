package main

import (
	"GoGit-Integration/pkg/gitapi"
	"GoGit-Integration/pkg/goGit"
	"fmt"
)

func main() {
	names := gitapi.GetList()

	fmt.Println("Repositories Paths:")
	for i := 0; i < len(names); i++ {
		fmt.Printf("%v. Avanis-GmbH/%v\n", i, names[i])
	}

	goGit.Clone(names)
}