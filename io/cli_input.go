package io

import (
	"bufio"
	"fmt"
	"os"
)

func AskUserMenuChoice() int {
	var choice int
	fmt.Print("Choose : ")
	fmt.Scan(&choice)
	return choice
}

func AskNewUserDetails() (string, int, string, int, string) {
	var name, address string
	var age, rollnumber int
	var courses_string string

	fmt.Print("full name : ")
	fmt.Scan(&name)

	fmt.Print("age : ")
	fmt.Scan(&age)

	fmt.Print("address : ")
	fmt.Scan(&address)

	fmt.Print("roll number : ")
	fmt.Scan(&rollnumber)

	fmt.Print("courses [A,B,C,D,E,F] (select atleast 4 space separated) : ")

	in := bufio.NewReader(os.Stdin)
	courses_string, err := in.ReadString('\n')
	// fmt.Scan(&courses_string)

	if err != nil {
		fmt.Println("error : reading the courses")
	}

	return name, age, address, rollnumber, courses_string
}

func AskUserChoiceForUserListDisplay() string {
	var sortPreference string
	fmt.Print("sort preference key [ name | age | address | roll ]: ")
	fmt.Scan(&sortPreference)
	return sortPreference
}

func AskUserRollNumberToDelete() int {
	var rollnumber int
	fmt.Print("roll number :")
	fmt.Scan(&rollnumber)
	return rollnumber
}

func AskUserToSaveUserDataOnDisk() string {
	var saveOnDisk string
	fmt.Print("save data on disk [y/n] :")
	fmt.Scan(&saveOnDisk)
	return saveOnDisk
}
