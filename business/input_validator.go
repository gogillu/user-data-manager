package business

import (
	"fmt"
	"problem2/user"
)

func validateNewUserDetails(usr user.User) bool {

	// check if the roll number already exists
	if _, alreadyExists := user.UserMap[usr.GetRollNumber()]; alreadyExists {
		fmt.Println("-*-*- User with roll number already exists, can not add user with given details, try again.")
		return false
	}

	// check if courses are >= 4
	if len(usr.GetCourses()) < 4 || len(usr.GetCourses()) > 6 {
		return false
	}
	return true
}
