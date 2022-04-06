package business

import (
	"fmt"
	"os"
	"strings"

	"problem2/io"
	"problem2/repository"
	"problem2/user"
	"problem2/user/course"
)

func Menu() {
	for {
		io.ShowMenuOptions()
		if takeActionBasedOnUserChoice(io.AskUserMenuChoice()) {
			break
		}
	}
}

func takeActionBasedOnUserChoice(userChoice int) bool {
	switch userChoice {
	case 1:
		askNewUserDetailsAndAdd()
	case 2:
		displayUserByUserSortKeyChoice()
	case 3:
		askUserAndDeleteUserByRollNumber()
	case 4:
		serializedUserDataAndSaveToDisk()
	case 5:
		askUserToSaveDataOnDiskBeforeExit()
	}
	return false
}

func frameCoursesFromString(courses_string string) ([]course.Course, error) {
	coursesString := strings.Fields(courses_string)
	var courseList []course.Course
	var courseMap map[string]bool = make(map[string]bool)

	for _, c := range coursesString {
		if _, ok := course.AllCourses[c]; !ok {
			fmt.Println(ok)
			return courseList, fmt.Errorf("%s error parsing the given course, invalid course : ", c)
		}

		if _, ok := courseMap[c]; ok {
			return courseList, fmt.Errorf("%s error courses are duplicate : ", c)
		}

		courseList = append(courseList, course.AllCourses[c])
		courseMap[c] = true
	}

	return courseList, nil
}

func askNewUserDetailsAndAdd() {
	name, age, address, rollnumber, course_string := io.AskNewUserDetails()
	courses, err := frameCoursesFromString(course_string)

	if err != nil {
		fmt.Println("--x-- : error : course entered are not valid try again ", err)
		return
	}

	newUser := user.User{}
	newUser.Init(name, age, address, rollnumber, courses)

	if validateNewUserDetails(newUser) {
		//update user roll number map
		user.UserMap[newUser.GetRollNumber()] = newUser

		//show new user added
		fmt.Println("\n==> new User Added ")
		io.SingleUserDisplay(newUser)
		fmt.Println("")
	} else {
		fmt.Println("--x-- : error : details entrered do not form a valid user")
	}
}

func askUserAndDeleteUserByRollNumber() {
	rollNumber := io.AskUserRollNumberToDelete()
	user.DeleteUserByRollNumber(rollNumber)
}

func displayUserByUserSortKeyChoice() {
	sortKey := io.AskUserChoiceForUserListDisplay()
	io.DisplayUserDetails(sortKey)
}

func serializedUserDataAndSaveToDisk() {
	repository.SaveUsersToDisk(user.UserMap)
}

func askUserToSaveDataOnDiskBeforeExit() {
	saveOnDiskBeforeExit := io.AskUserToSaveUserDataOnDisk()
	if saveOnDiskBeforeExit == "y" {
		serializedUserDataAndSaveToDisk()
	}
	os.Exit(0)
}
