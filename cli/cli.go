package cli

import (
	"bufio"
	"fmt"
	"os"

	"gihub.com/gogillu/user-manager/user"
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

func ShowMenu() {
	fmt.Println()
	fmt.Println(" __________________")
	fmt.Println("|    User Menu     |")
	fmt.Println("|  -------------   |")
	fmt.Println("|   1. Add         |")
	fmt.Println("|   2. Display     |")
	fmt.Println("|   3. Delete      |")
	fmt.Println("|   4. Save        |")
	fmt.Println("|   5. Exit        |")
	fmt.Println("|__________________|")
	fmt.Println()
}

func DisplayUsers(users []user.User) {
	fmt.Println()
	fmt.Println("+-----------+--------------+----------+----------------+---------------------+")
	fmt.Printf("|%10s |%13s |%9s |%15s |%20s |\n", "Name", "Roll Number", "Age", "Address", "Courses")
	fmt.Println("+-----------+--------------+----------+----------------+---------------------+")
	for _, usr := range users {
		DisplayUser(usr)
	}
}

func DisplayUser(usr user.User) {
	var courses string
	for _, course := range usr.GetCourses() {
		courses += course.String() + ", "
	}

	fmt.Printf("|%10s |%13d |%9d |%15s |%20s |\n", usr.GetName(), usr.GetRollNumber(), usr.GetAge(), usr.GetAddress(), courses)
	fmt.Println("+-----------+--------------+----------+----------------+---------------------+")
}
