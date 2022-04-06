package repository

import (
	"io/ioutil"
	"problem2/user"
)

const User_Data_File_Location = "./data/users.json"

func readUsersDataFromDisk() string {

	file, error := ioutil.ReadFile(User_Data_File_Location)
	if error != nil {
		panic(error)
	}

	return string(file)
}

func GetUsersFromDisk() map[string][]user.User {
	return user.ConvertJsonStringToUserDetailsMap(readUsersDataFromDisk())
}
