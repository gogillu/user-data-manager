package app

import (
	"fmt"
	"strings"

	"gihub.com/gogillu/user-manager/cli"
	"gihub.com/gogillu/user-manager/repository"
	"gihub.com/gogillu/user-manager/user"
	"gihub.com/gogillu/user-manager/user/enum"
)

const (
	Datastore = "./data/users.json"
	Yes       = "y"
)

func Start() error {
	userRepo := repository.NewRepo()
	err := userRepo.Load(Datastore)
	if err != nil {
		return nil
	}

	exit := false
	for !exit {
		cli.ShowMenu()
		choice := cli.GetMenuChoice()
		var err error

		switch choice {
		case 1:
			err = AddUser(userRepo)
		case 2:
			err = ListUsers(userRepo)
		case 3:
			err = DeleteUser(userRepo)
		case 4:
			err = SaveUsers(userRepo)
		case 5:
			err = Exit(userRepo)
			exit = true
		default:
			continue
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func getCourses(courses string) ([]enum.Course, error) {
	coursesList := strings.Fields(courses)
	var courseList []enum.Course

	for _, c := range coursesList {
		course, err := enum.CourseString(c)
		if err != nil {
			return courseList, fmt.Errorf("%s error parsing the provided course, invalid course : ", c)
		}

		courseList = append(courseList, course)
	}

	return courseList, nil
}

func AddUser(userRepo repository.Repository) error {
	name, age, address, rollnumber, course_string := cli.GetUser()
	courses, err1 := getCourses(course_string)
	if err1 != nil {
		fmt.Println("error : course entered are not valid try adding again", err1)
		return nil
	}

	usr, err2 := user.New(name, age, address, rollnumber, courses)
	if err2 != nil {
		fmt.Println("error : ", err2, "try again!")
		return nil
	}

	err3 := userRepo.Add(usr)
	if err3 != nil {
		fmt.Println("error : ", err3, "try again!")
		return nil
	}

	fmt.Println("\n new user added")
	cli.DisplayUser(usr)
	return nil
}

func ListUsers(userRepo repository.Repository) error {
	sortKey := cli.GetListFilter()
	users, err := userRepo.List(sortKey)
	if err != nil {
		return fmt.Errorf("error listing users %v", err)
	}

	cli.DisplayUsers(users)
	return nil
}

func DeleteUser(userRepo repository.Repository) error {
	rollNo := cli.GetRollNo()
	err := userRepo.Delete(rollNo)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("user deleted successfully with roll number ", rollNo)
	return nil
}

func SaveUsers(userRepo repository.Repository) error {
	err := userRepo.Save()
	if err != nil {
		return err
	}

	fmt.Println("Saved on disk")
	return nil
}

func Exit(repo repository.Repository) error {
	choice := cli.ConfirmSave()
	if choice == Yes {
		if err := SaveUsers(repo); err != nil {
			return err
		}
	}
	return nil
}
