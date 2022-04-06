package user

import (
	"encoding/json"
	"fmt"
	"sort"
)

// var Users []User
var UserMap map[int]User = make(map[int]User)

func InitializeUserData(savedUsers []User) {
	// Users = savedUsers

	for _, usr := range savedUsers {
		UserMap[usr.GetRollNumber()] = usr
	}
}

func DeleteUserByRollNumber(rollNumber int) bool {
	if _, userExists := UserMap[rollNumber]; userExists {
		fmt.Println("===> Found a user in repository, deleting...")

		// remove user from
		delete(UserMap, rollNumber)
		fmt.Println("===> Deleted...")
	} else {
		fmt.Println("---> No user exists with given roll number")
	}
	return true
}

func GetUserListFromUserMap(usrMap map[int]User) []User {
	var users []User

	for _, usr := range usrMap {
		users = append(users, usr)
	}

	return users
}

func SerializeUsersData(users []User) string {
	serializedUserData, err := json.Marshal(users)

	if err != nil {
		panic(err)
	}

	return string(serializedUserData)
}

func sortUsersByName(users []User) []User {
	sort.Slice(users, func(i, j int) bool {
		if users[i].Name < users[j].Name {
			return true
		} else if users[i].Name == users[j].Name {
			if users[i].RollNumber < users[j].RollNumber {
				return true
			}
		}
		return false
	})
	return users
}

func sortUsersByAge(users []User) []User {
	sort.Slice(users, func(i, j int) bool {
		if users[i].Age < users[j].Age {
			return true
		} else if users[i].Age == users[j].Age {
			if users[i].RollNumber < users[j].RollNumber {
				return true
			}
		}
		return false
	})
	return users
}

func sortUsersByAddress(users []User) []User {
	sort.Slice(users, func(i, j int) bool {
		if users[i].Address < users[j].Address {
			return true
		} else if users[i].Address == users[j].Address {
			if users[i].RollNumber < users[j].RollNumber {
				return true
			}
		}
		return false
	})
	return users
}

func sortUsersByRollNumber(users []User) []User {
	sort.Slice(users, func(i, j int) bool {
		if users[i].RollNumber < users[j].RollNumber {
			return true
		}
		return false
	})
	return users
}

func SortUsersByKey(users []User, sortKey string) []User {
	switch sortKey {
	case "name":
		fmt.Println("User List sorted by : ", sortKey)
		return sortUsersByName(users)
	case "age":
		fmt.Println("User List sorted by : ", sortKey)
		return sortUsersByAge(users)
	case "address":
		fmt.Println("User List sorted by : ", sortKey)
		return sortUsersByAddress(users)
	default:
		fmt.Println("User List sorted by : Roll Number")
		return sortUsersByRollNumber(users)
	}
}
