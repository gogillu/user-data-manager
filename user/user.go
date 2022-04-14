package user

import (
	"encoding/json"
	"fmt"

	"gihub.com/gogillu/user-manager/user/enum"

	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	MinCourses = 4
	MaxCourses = 6
)

type User struct {
	Name       string        `json:"name"`
	Age        int           `json:"age"`
	Address    string        `json:"address"`
	RollNumber int           `json:"roll_number"`
	Courses    []enum.Course `json:"courses"`
}

func New(name string, age int, address string, rollnumber int, courses []enum.Course) (User, error) {
	var usr User
	usr.Name = name
	usr.Age = age
	usr.Address = address
	usr.RollNumber = rollnumber
	usr.Courses = courses

	if err := usr.validate(); err != nil {
		return User{}, fmt.Errorf("error initizliating User %v", err)
	}

	return usr, nil
}

func (usr User) validate() error {
	return validation.ValidateStruct(&usr,
		validation.Field(&usr.RollNumber, validation.Required),
		validation.Field(&usr.Courses, validation.Required, validation.By(validateCourses)),
	)
}

// validate Courses (1) Should be 4 to 6 (2) No repeatation
func validateCourses(value interface{}) error {
	courses := value.([]enum.Course)
	if len(courses) < MinCourses || len(courses) > MaxCourses {
		return fmt.Errorf("error : found invalid number of courses %d", len(courses))
	}

	var courseMap map[enum.Course]bool = make(map[enum.Course]bool)
	for _, c := range courses {
		if _, duplicate := courseMap[c]; duplicate {
			return fmt.Errorf("error : courses are duplicate %v", c)
		}
		courseMap[c] = true
	}

	return nil
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

func (u User) GetCourses() []enum.Course {
	return u.Courses
}

func (u User) GetCoursesString() string {
	var courses string

	for _, course := range u.GetCourses() {
		courses += course.String() + ", "
	}
	return courses
}

func EncodeUsers(users []User) (string, error) {
	serializedUserData, err := json.Marshal(users)

	if err != nil {
		return "", err
	}

	return string(serializedUserData), nil
}

func DecodeUsers(userDetailsString string) ([]User, error) {

	var users []User
	err := json.Unmarshal([]byte(userDetailsString), &users)

	if err != nil {
		return []User{}, fmt.Errorf("error in decoding users %v", err)
	}

	return users, nil
}
