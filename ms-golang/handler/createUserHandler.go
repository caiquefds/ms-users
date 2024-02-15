package handler

import (
	"errors"

	"github.com/caiquefds/ms-users/model"
	"github.com/google/uuid"
)

func CreateUser(createUserRequest *model.CreateUserRequest) (userId string, errorResponse error) {
	logger.Info("Creating new user...")

	user := &model.User{
		Id:       uuid.New().String(),
		Password: createUserRequest.Password,
		Name:     createUserRequest.Name,
		Email:    createUserRequest.Email,
		Status:   "ENABLED",
	}

	if err := db.Create(&user).Error; err != nil {
		logger.ErrF("Failed to create user: %+v", err.Error())
		err = errors.New("422.001")
		return "", err
	}

	logger.Info("User created!")
	return user.Id, nil
}
