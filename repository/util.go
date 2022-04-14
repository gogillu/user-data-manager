package repository

import (
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"user-manager/user"
)

func createEmptyFile(filepath string) error {
	if _, err := os.Stat(filepath); err != nil {
		_, e := os.Create(filepath)
		if e != nil {
			return e
		}
	}

	err2 := WriteDisk(filepath, "[]")
	if err2 != nil {
		return err2
	}

	return nil
}

func readDisk(file string) (string, error) {

	fileData, err := ioutil.ReadFile(file)
	if err != nil {
		err := createEmptyFile(file)
		return "[]", err
	}

	return string(fileData), nil
}

func WriteDisk(filepath string, userDetails string) error {
	err := ioutil.WriteFile(filepath, []byte(userDetails), 0644)
	if err != nil {
		return err
	}

	return nil
}

func SortUsers(users []user.User, key string) []user.User {
	sort.SliceStable(users, func(j, i int) bool {
		var cmp int
		switch key {
		case Name:
			cmp = strings.Compare(users[i].GetName(), users[j].GetName())
		case Address:
			cmp = strings.Compare(users[i].GetAddress(), users[j].GetAddress())
		case Age:
			cmp = users[i].GetAge() - users[j].GetAge()
		}

		if cmp == 0 {
			return users[i].GetRollNumber() > users[j].GetRollNumber()
		}
		return cmp > 0
	})
	return users
}
