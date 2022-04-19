package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"gihub.com/gogillu/user-manager/user"
)

const (
	Name    = "name"
	Age     = "age"
	Address = "address"
)

type Repository interface {
	Load(filepath string) error
	List(sortKey string) ([]user.User, error)
	Add(usr user.User) error
	Delete(rollno int) error
	Save() error
}

type repository struct {
	users map[int]user.User
	file  string
}

func NewRepo() *repository {
	return &repository{}
}

func (repo *repository) Load(filepath string) error {
	fileData, err := readDisk(filepath)
	if err != nil {
		return fmt.Errorf("error reading file from disk %v", err)
	}

	users, err := DecodeUsers(fileData)
	if err != nil {
		return fmt.Errorf("error decoding users to list %v", err)
	}

	repo.file = filepath
	repo.users = make(map[int]user.User)
	for _, user := range users {
		repo.users[user.GetRollNumber()] = user
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

func (repo *repository) List(key string) ([]user.User, error) {
	sortedUsers := SortUsers(repo.users, key)
	return sortedUsers, nil
}

func SortUsers(usersMap map[int]user.User, key string) []user.User {

	var users []user.User

	for _, usr := range usersMap {
		users = append(users, usr)
	}

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

func (repo *repository) Add(usr user.User) error {
	if _, exists := repo.users[usr.GetRollNumber()]; exists {
		return fmt.Errorf("error : user with %d roll number already exists", usr.GetRollNumber())
	}

	repo.users[usr.GetRollNumber()] = usr
	return nil
}

func (repo *repository) Delete(rollNo int) error {
	if _, userExists := repo.users[rollNo]; userExists {
		delete(repo.users, rollNo)
		return nil
	}
	return fmt.Errorf("error deleting user with rollno, user doesn't exist %d", rollNo)
}

func (repo *repository) Save() error {
	encodedUsers, err := EncodeUsers(repo.users)
	if err != nil {
		return err
	}

	err = WriteDisk(repo.file, encodedUsers)
	if err != nil {
		return err
	}

	return nil
}

func WriteDisk(filepath string, userDetails string) error {
	err := ioutil.WriteFile(filepath, []byte(userDetails), 0644)
	if err != nil {
		return err
	}

	return nil
}

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

func EncodeUsers(usersMap map[int]user.User) (string, error) {
	serializedUserData, err := json.Marshal(usersMap)

	if err != nil {
		return "", fmt.Errorf("error encoding users %v", err)
	}

	return string(serializedUserData), nil
}

func DecodeUsers(userDetailsString string) (map[int]user.User, error) {
	var usersMap map[int]user.User
	err := json.Unmarshal([]byte(userDetailsString), &usersMap)

	if err != nil {
		return map[int]user.User{}, fmt.Errorf("error in decoding users %v", err)
	}

	return usersMap, nil
}
