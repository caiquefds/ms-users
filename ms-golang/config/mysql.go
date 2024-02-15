package config

import (
	"github.com/caiquefds/ms-users/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMysqlDb() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	return db, err
}
