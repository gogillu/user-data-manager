package user

import (
	"encoding/json"
	"problem2/user/course"
)

type User struct {
	Name       string
	Age        int
	Address    string
	RollNumber int
	Courses    []course.Course
}

func (u *User) Init(name string, age int, address string, rollnumber int, courses []course.Course) {
	u.Name = name
	u.Age = age
	u.Address = address
	u.RollNumber = rollnumber
	u.Courses = courses
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetRollNumber() int {
	return u.RollNumber
}

func (u User) GetAge() int {
	return u.Age
}

func (u User) GetAddress() string {
	return u.Address
}

func (u User) GetCourses() []course.Course {
	return u.Courses
}

func (u User) GetCoursesString() string {
	var courses string

	for _, c := range u.Courses {
		courses = courses + string(c) + ", "
	}

	return courses
}

func ConvertJsonStringToUserDetailsMap(userDetailsString string) map[string][]User {

	userDetailsMap := make(map[string][]User)
	err := json.Unmarshal([]byte(userDetailsString), &userDetailsMap)

	if err != nil {
		panic(err)
	}

	// fmt.Println(userDetailsMap)
	return userDetailsMap
}
