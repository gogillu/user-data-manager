package main

import (
	"fmt"

	"gihub.com/gogillu/user-manager/app"
)

func main() {
	err := app.Start()
	if err != nil {
		fmt.Println(err)
	}
}
