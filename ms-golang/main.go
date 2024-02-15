package main

import (
	"github.com/caiquefds/ms-users/config"
	route "github.com/caiquefds/ms-users/routes"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.InitializeConfigs()
	// validation.InitializeValidations()
	if err != nil {
		logger.ErrF("Config initialization error: v%", err)
		return
	}
	route.Initialize()
}
