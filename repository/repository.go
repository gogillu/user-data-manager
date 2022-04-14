package repository

import (
	"fmt"
	"user-manager/user"
)

const (
	Name    = "name"
	Age     = "age"
	Address = "address"
)

type Repository interface {
	Load(string) error
	List(string) ([]user.User, error)
	Add(usr user.User) error
	Delete(int) error
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

	users, err := user.DecodeUsers(fileData)
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

func (repo *repository) List(key string) ([]user.User, error) {
	var users []user.User
	for _, usr := range repo.users {
		users = append(users, usr)
	}

	sortedUsers := SortUsers(users, key)
	return sortedUsers, nil
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
	var users []user.User
	for _, usr := range repo.users {
		users = append(users, usr)
	}

	encodedUsers, err := user.EncodeUsers(users)
	if err != nil {
		return err
	}

	err2 := WriteDisk(repo.file, encodedUsers)
	if err2 != nil {
		return err
	}

	return nil
}
