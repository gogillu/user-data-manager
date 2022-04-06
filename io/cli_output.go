package io

import (
	"fmt"
	"problem2/user"
)

func ShowMenuOptions() {
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

func DisplayUserDetails(sortKey string) {
	users := user.SortUsersByKey(user.GetUserListFromUserMap(user.UserMap), sortKey)
	fmt.Println()
	fmt.Println("+-----------+--------------+----------+----------------+---------------------+")
	fmt.Printf("|%10s |%13s |%9s |%15s |%20s |\n", "Name", "Roll Number", "Age", "Address", "Courses")
	for _, usr := range users {
		// fmt.Print(usr)
		fmt.Println("+-----------+--------------+----------+----------------+---------------------+")
		fmt.Printf("|%10s |%13d |%9d |%15s |%20s |\n", usr.GetName(), usr.GetRollNumber(), usr.GetAge(), usr.GetAddress(), usr.GetCoursesString())
	}
	fmt.Println("+-----------+--------------+----------+----------------+---------------------+")
}

func SingleUserDisplay(singleuser user.User) {
	fmt.Println("+-----------+--------------+----------+----------------+---------------------+")
	fmt.Printf("|%10s |%13d |%9d |%15s |%20s |\n", singleuser.GetName(), singleuser.GetRollNumber(), singleuser.GetAge(), singleuser.GetAddress(), singleuser.GetCoursesString())
	fmt.Println("+-----------+--------------+----------+----------------+---------------------+")
}
