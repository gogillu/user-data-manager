package bootstrap

import (
	"problem2/repository"
	"problem2/user"
)

func Load() {
	users := repository.GetUsersFromDisk()
	user.InitializeUserData(users["users"])
}
