package main

import (
	"fmt"
	"user-manager/app"
)

func main() {
	err := app.Start()
	if err != nil {
		fmt.Println(err)
	}
}
