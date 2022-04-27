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
	DatastoreFilepath = "./data/users.json"
	Yes               = "y"
	OptionAdd         = 1
	OptionList        = 2
	OptionDelete      = 3
	OptionSave        = 4
	OptionExit        = 5
)

func Start() error {
	userRepo := repository.NewRepo()
	err := userRepo.Load(DatastoreFilepath)
	if err != nil {
		return err
	}

	exit := false
	for !exit {
		cli.ShowMenu()
		choice := cli.GetMenuChoice()

		switch choice {
		case OptionAdd:
			err = AddUser(userRepo)
		case OptionList:
			err = ListUsers(userRepo)
		case OptionDelete:
			err = DeleteUser(userRepo)
		case OptionSave:
			err = SaveUsers(userRepo)
		case OptionExit:
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
	obtainedCourseList := strings.Fields(courses)
	var framedCourseList []enum.Course

	for _, c := range obtainedCourseList {
		course, err := enum.CourseString(c)
		if err != nil {
			return framedCourseList, fmt.Errorf("%s error parsing the provided course, invalid course : ", c)
		}

		framedCourseList = append(framedCourseList, course)
	}

	return framedCourseList, nil
}

func AddUser(userRepo repository.Repository) error {
	name, age, address, rollnumber, course_string := cli.GetUser()
	courses, err := getCourses(course_string)
	if err != nil {
		fmt.Println("error : course entered are not valid try adding again", err)
		return nil
	}

	var usr user.User
	usr, err = user.New(name, age, address, rollnumber, courses)
	if err != nil {
		fmt.Println("error : ", err, "try again!")
		return nil
	}

	err = userRepo.Add(usr)
	if err != nil {
		fmt.Println("error : ", err, "try again!")
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
