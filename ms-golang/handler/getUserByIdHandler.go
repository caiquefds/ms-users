package handler

import (
	"errors"

	"github.com/caiquefds/ms-users/model"
)

func GetUserById(userId string) (*model.GetUserByIdResponse, error) {
	logger.InfoF("Retrivieng user by id %s...", userId)

	user := &model.User{}

	if err := db.Find(user, "id = ?", userId).Error; err != nil {
		logger.ErrF("User not found: %+v", err.Error())
		return nil, errors.New("User not found!")
	}

	response := model.GetUserByIdResponse{
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
		Status:   user.Status,
	}

	return &response, nil
}
