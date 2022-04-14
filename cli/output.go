package cli

import (
	"fmt"
	"user-manager/user"
)

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

func DisplayUser(singleuser user.User) {
	fmt.Printf("|%10s |%13d |%9d |%15s |%20s |\n", singleuser.GetName(), singleuser.GetRollNumber(), singleuser.GetAge(), singleuser.GetAddress(), singleuser.GetCoursesString())
	fmt.Println("+-----------+--------------+----------+----------------+---------------------+")
}
