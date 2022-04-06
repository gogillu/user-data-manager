package user

import (
	"problem2/user/course"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertJsonStringToUserDetailsMap(t *testing.T) {

	jsonString := `{"users":[{ "name" : "Ram" , "age" : 15 , "address" : "ABC", "rollnumber" : 1234 , "courses" : [ "A" , "B" , "C" , "E" ] }]}`

	user1 := User{Name: "Ram", Age: 15, Address: "ABC", RollNumber: 1234, Courses: []course.Course{"A", "B", "C", "E"}}
	expectedUserDetailMap := map[string][]User{"users": {user1}}

	actualOutput := ConvertJsonStringToUserDetailsMap(jsonString)

	assert.Equal(t, actualOutput, expectedUserDetailMap, "comparing text")
}
