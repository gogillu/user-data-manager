package cli

import (
	"bufio"
	"fmt"
	"os"
)

func GetMenuChoice() int {
	var choice int
	fmt.Print("Choose : ")
	fmt.Scan(&choice)
	return choice
}

func GetUser() (string, int, string, int, string) {
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
	if err != nil {
		fmt.Println("error : reading the courses")
	}

	return name, age, address, rollnumber, courses_string
}

func GetListFilter() string {
	var sortPreference string
	fmt.Print("sort preference key [ name | age | address | roll ]: ")
	fmt.Scan(&sortPreference)
	return sortPreference
}

func GetRollNo() int {
	var rollnumber int
	fmt.Print("roll number :")
	fmt.Scan(&rollnumber)
	return rollnumber
}

func ConfirmSave() string {
	var saveOnDisk string
	fmt.Print("save data on disk [y/n] :")
	fmt.Scan(&saveOnDisk)
	return saveOnDisk
}
