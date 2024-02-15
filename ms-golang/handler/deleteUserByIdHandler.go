package handler

import (
	"github.com/caiquefds/ms-users/model"
)

func DeleteUserById(userId string) {
	logger.InfoF("Deleting user by id %s...", userId)

	if err := db.Where("id = ?", userId).Delete(&model.User{}).Error; err != nil {
		logger.ErrF("User not found: %+v", err.Error())
	} else {
		logger.Info("User deleted.")
	}

}
