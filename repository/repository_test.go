package repository

import (
	"fmt"
	"testing"
	"user-manager/user"
	"user-manager/user/enum"

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
	testRepo.Load(Datastore)

	test := []struct {
		usr    user.User
		expErr error
	}{
		{
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
			assert.Equal(t, actualErr, tc.expErr, "Compare err")
			assert.Equal(t, tc.usr, testRepo.users[tc.usr.GetRollNumber()], "test user in map")
		} else {
			assert.NotEqual(t, actualErr, nil, "Compare err")
			assert.NotEqual(t, tc.usr, testRepo.users[tc.usr.GetRollNumber()], "test user in map")
		}

	}
}
