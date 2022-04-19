package repository

import (
	"fmt"
	"testing"

	"gihub.com/gogillu/user-manager/user"
	"gihub.com/gogillu/user-manager/user/enum"

	"github.com/stretchr/testify/assert"
)

const (
	Datastore = "../data/users_test.json"
)

func TestLoad(t *testing.T) {

	testRepo := NewRepo()

	test := struct {
		filepath           string
		expErr             error
		expLoadedRepoUsers map[int]user.User
	}{
		Datastore,
		nil,
		map[int]user.User{},
	}

	actualErr := testRepo.Load(Datastore)

	assert.Equal(t, actualErr, test.expErr, "Test Load")
	assert.Equal(t, test.expLoadedRepoUsers, testRepo.users, "Test Load - empty user map")

}

func TestAdd(t *testing.T) {
	testRepo := NewRepo()
	_ = testRepo.Load(Datastore)

	test := []struct {
		name   string
		usr    user.User
		expErr error
	}{
		{
			"Valid User 1",
			user.User{
				Name:       "Ram",
				Age:        11,
				Address:    "Lane-1",
				RollNumber: 101,
				Courses:    []enum.Course{enum.A, enum.B, enum.C, enum.D},
			},
			nil,
		},
		{
			"Valid user 2",
			user.User{
				Name:       "Shyam",
				Age:        12,
				Address:    "Lane-2",
				RollNumber: 102,
				Courses:    []enum.Course{enum.A, enum.B, enum.C, enum.E, enum.F},
			},
			nil,
		},
		{
			"Valid user 3 with duplicate roll number",
			user.User{
				Name:       "Vivek",
				Age:        40,
				Address:    "Lane-9",
				RollNumber: 102,
				Courses:    []enum.Course{enum.A, enum.B, enum.C, enum.E},
			},
			fmt.Errorf("error : roll number already exists"),
		},
	}

	assert.Equal(t, len(testRepo.users), 0, "Base case")

	for _, tc := range test {
		actualErr := testRepo.Add(tc.usr)

		if tc.expErr == nil {
			assert.Equal(t, actualErr, tc.expErr, tc.name)
			assert.Equal(t, tc.usr, testRepo.users[tc.usr.GetRollNumber()], tc.name)
		} else {
			assert.NotEqual(t, actualErr, nil, tc.name)
			assert.NotEqual(t, tc.usr, testRepo.users[tc.usr.GetRollNumber()], tc.name)
		}

	}
}

func TestDelete(t *testing.T) {

	testRepo := NewRepo()
	_ = testRepo.Load(Datastore)

	initialData := []user.User{
		{
			Name:       "Ram",
			Age:        11,
			Address:    "Lane-1",
			RollNumber: 101,
			Courses:    []enum.Course{enum.A, enum.B, enum.C, enum.D},
		},
		{
			Name:       "Shyam",
			Age:        12,
			Address:    "Lane-2",
			RollNumber: 102,
			Courses:    []enum.Course{enum.A, enum.B, enum.C, enum.E, enum.F},
		},
		{
			Name:       "Manoj",
			Age:        13,
			Address:    "Lane-3",
			RollNumber: 103,
			Courses:    []enum.Course{enum.B, enum.C, enum.E, enum.F},
		},
	}

	for _, usr := range initialData {
		_ = testRepo.Add(usr)
	}

	test := []struct {
		name   string
		rollNo int
		expErr error
	}{
		{
			"Deleting a valid roll number",
			101,
			nil,
		},
		{
			"Deleting a invalid roll number",
			106,
			fmt.Errorf("error deleting user with rollno, user doesn't exist 106"),
		},
	}

	for _, tc := range test {
		actualErr := testRepo.Delete(tc.rollNo)

		fmt.Println(testRepo)

		if tc.expErr == nil {
			assert.Equal(t, actualErr, tc.expErr, tc.name)
			assert.Equal(t, user.User{}, testRepo.users[tc.rollNo], tc.name)
		} else {
			assert.Equal(t, actualErr, tc.expErr, tc.name)
			assert.Equal(t, user.User{}, testRepo.users[tc.rollNo], tc.name)
		}

	}
}

func TestEncodeUsers(t *testing.T) {

	tests := []struct {
		userMap map[int]user.User
		expOut  string
		expErr  error
	}{
		{
			map[int]user.User{
				101: {
					Name:       "Ram",
					Age:        11,
					Address:    "Lane-1",
					RollNumber: 101,
					Courses:    []enum.Course{enum.A, enum.B, enum.C, enum.D},
				},
				102: {
					Name:       "Shyam",
					Age:        12,
					Address:    "Lane-2",
					RollNumber: 102,
					Courses:    []enum.Course{enum.A, enum.B, enum.C, enum.D, enum.E},
				},
			},
			`{"101":{"name":"Ram","age":11,"address":"Lane-1","roll_number":101,"courses":[0,1,2,3]},"102":{"name":"Shyam","age":12,"address":"Lane-2","roll_number":102,"courses":[0,1,2,3,4]}}`,
			nil,
		},
	}

	for _, tests := range tests {
		actualOut, actualErr := EncodeUsers(tests.userMap)
		assert.Equal(t, actualOut, tests.expOut, "Test User Encoding")
		assert.Equal(t, actualErr, tests.expErr, "Test User Encoding")
	}
}

func TestDecodeUsers(t *testing.T) {

	tests := []struct {
		name   string
		users  string
		expOut map[int]user.User
		expErr error
	}{
		{
			"Valid Serialized user string to userMap test",
			`{"101":{"name":"Ram","age":11,"address":"Lane-1","roll_number":101,"courses":[0,1,2,3]},"102":{"name":"Shyam","age":12,"address":"Lane-2","roll_number":102,"courses":[0,1,2,3,4]}}`,
			map[int]user.User{
				101: {
					Name:       "Ram",
					Age:        11,
					Address:    "Lane-1",
					RollNumber: 101,
					Courses:    []enum.Course{enum.A, enum.B, enum.C, enum.D},
				},
				102: {
					Name:       "Shyam",
					Age:        12,
					Address:    "Lane-2",
					RollNumber: 102,
					Courses:    []enum.Course{enum.A, enum.B, enum.C, enum.D, enum.E},
				},
			},
			nil,
		},
		{
			"Invalid Serialized user string to userMap test",
			`{"101:{"name":"Ram","age":11,"address":"Lane-1","roll_number":101,"courses":[0,1,2,3]},"102":{"name":"Shyam","age":12,"address":"Lane-2","roll_number":102,"courses":[0,1,2,3,4]}}]`,
			map[int]user.User{
				101: {
					Name:       "Ram",
					Age:        11,
					Address:    "Lane-1",
					RollNumber: 101,
					Courses:    []enum.Course{enum.A, enum.B, enum.C, enum.D},
				},
				102: {
					Name:       "Shyam",
					Age:        12,
					Address:    "Lane-2",
					RollNumber: 102,
					Courses:    []enum.Course{enum.A, enum.B, enum.C, enum.D, enum.E},
				},
			},
			fmt.Errorf("error decoding invalid string input"),
		},
	}

	for _, tests := range tests {
		actualOut, actualErr := DecodeUsers(tests.users)
		if tests.expErr == nil {
			assert.Equal(t, actualOut, tests.expOut, tests.name)
			assert.Equal(t, actualErr, tests.expErr, tests.name)
		} else {
			assert.NotEqual(t, actualErr, tests.expErr, tests.name)
		}
	}
}
