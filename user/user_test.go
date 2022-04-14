package user

import (
	"fmt"
	"testing"

	"gihub.com/gogillu/user-manager/user/enum"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	tests := []struct {
		usr    User
		expErr error
	}{
		{
			User{
				"Ram",
				11,
				"Lane-1",
				101,
				[]enum.Course{enum.A, enum.B, enum.C, enum.D},
			},
			nil,
		},
		{
			User{
				"Shyam",
				12,
				"Lane-2",
				102,
				[]enum.Course{enum.A, enum.B, enum.C},
			},
			fmt.Errorf("error : found invalid number of courses"),
		},
		{
			User{
				"Manoj",
				13,
				"Lane-3",
				103,
				[]enum.Course{enum.A, enum.B, enum.C, enum.C},
			},
			fmt.Errorf("error : courses are duplicate"),
		},
	}

	for _, tc := range tests {
		actualUser, actualErr := New(tc.usr.Name, tc.usr.Age, tc.usr.Address, tc.usr.RollNumber, tc.usr.Courses)
		if tc.expErr == nil {
			assert.Equal(t, actualErr, tc.expErr, "Test New User")
			assert.Equal(t, actualUser, tc.usr, "Test New User")
		} else {
			assert.NotEqual(t, actualErr, tc.expErr, "Test New User")
			assert.NotEqual(t, actualUser, tc.usr, "Test New User")
		}
	}
}

func TestValidate(t *testing.T) {

	tests := []struct {
		usr    User
		expErr error
	}{
		{
			User{
				"Ram",
				11,
				"Lane-1",
				101,
				[]enum.Course{enum.A, enum.B, enum.C, enum.D},
			},
			nil,
		},
		{
			User{
				"Shyam",
				12,
				"Lane-2",
				102,
				[]enum.Course{enum.A, enum.B, enum.C},
			},
			fmt.Errorf("error : found invalid number of courses"),
		},
		{
			User{
				"Manoj",
				13,
				"Lane-3",
				103,
				[]enum.Course{enum.A, enum.B, enum.C, enum.C},
			},
			fmt.Errorf("error : courses are duplicate"),
		},
	}

	for _, tests := range tests {
		actualErr := tests.usr.validate()
		if tests.expErr == nil {
			assert.Equal(t, actualErr, tests.expErr, "Test Course Validation - Positive")
		} else {
			assert.NotEqual(t, actualErr, tests.expErr, "Test Course Validation - Negative")
		}
	}
}

func TestEncodeUsers(t *testing.T) {

	tests := []struct {
		users  []User
		expOut string
		expErr error
	}{
		{
			[]User{
				{
					"Ram",
					11,
					"Lane-1",
					101,
					[]enum.Course{enum.A, enum.B, enum.C, enum.D},
				},
				{
					"Shyam",
					12,
					"Lane-2",
					102,
					[]enum.Course{enum.A, enum.B, enum.C, enum.D, enum.F},
				},
			},
			`[{"name":"Ram","age":11,"address":"Lane-1","roll_number":101,"courses":[0,1,2,3]},{"name":"Shyam","age":12,"address":"Lane-2","roll_number":102,"courses":[0,1,2,3,5]}]`,
			nil,
		},
	}

	for _, tests := range tests {
		actualOut, actualErr := EncodeUsers(tests.users)
		assert.Equal(t, actualOut, tests.expOut, "Test User Encoding")
		assert.Equal(t, actualErr, tests.expErr, "Test User Encoding")
	}
}

func TestDecodeUsers(t *testing.T) {

	tests := []struct {
		users  string
		expOut []User
		expErr error
	}{
		{
			`[{"name":"Ram","age":11,"address":"Lane-1","roll_number":101,"courses":[0,1,2,3]},{"name":"Shyam","age":12,"address":"Lane-2","roll_number":102,"courses":[0,1,2,3,5]}]`,
			[]User{
				{
					"Ram",
					11,
					"Lane-1",
					101,
					[]enum.Course{enum.A, enum.B, enum.C, enum.D},
				},
				{
					"Shyam",
					12,
					"Lane-2",
					102,
					[]enum.Course{enum.A, enum.B, enum.C, enum.D, enum.F},
				},
			},
			nil,
		},
		{
			`[{"name":"Ram","age":11,"address":"Lane-1","roll_number":101,"courses":[0,1,2,3]},{"name":"Shyam","age":12,"address":"Lane-2","roll_number":102,"courses":[0,1,2,3,5]}`,
			[]User{
				{
					"Ram",
					11,
					"Lane-1",
					101,
					[]enum.Course{enum.A, enum.B, enum.C, enum.D},
				},
				{
					"Shyam",
					12,
					"Lane-2",
					102,
					[]enum.Course{enum.A, enum.B, enum.C, enum.D, enum.F},
				},
			},
			fmt.Errorf("error decoding invalid string input"),
		},
	}

	for _, tests := range tests {
		actualOut, actualErr := DecodeUsers(tests.users)
		if tests.expErr == nil {
			assert.Equal(t, actualOut, tests.expOut, "Test User Decoding Error")
			assert.Equal(t, actualErr, tests.expErr, "Test User Decoding Value")
		} else {
			assert.NotEqual(t, actualErr, tests.expErr, "Test User Decoding Value Invalid String")
		}
	}
}
