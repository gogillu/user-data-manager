package repository

import (
	"fmt"
	"io/ioutil"
	"problem2/user"
)

func WriteUserDetailsToDisk(userDetails string) bool {

	err := ioutil.WriteFile(User_Data_File_Location, []byte(`{"users":`+userDetails+`}`), 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data Saved On Disk")
	return true
}

func SaveUsersToDisk(users map[int]user.User) bool {
	WriteUserDetailsToDisk(user.SerializeUsersData(user.GetUserListFromUserMap(users)))
	return true
}
