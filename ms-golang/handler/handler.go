package handler

import (
	"github.com/caiquefds/ms-users/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB

)

func Initialize() {
	logger = config.GetLogger("UsersService")
	db = config.GetMysqlDb()
}
