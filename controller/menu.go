package controller

import (
	"problem2/bootstrap"
	"problem2/business"
)

func Menu() {
	bootstrap.Load()
	business.Menu()
}
