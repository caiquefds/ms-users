package handler

import (
	"reflect"

	"github.com/caiquefds/ms-users/model"
)

func UpdateUser(updateUserRequest *model.UpdateUserRequest) {
	logger.Info("Updating user...")

	user := &model.User{
		Id:       updateUserRequest.Id,
		Password: updateUserRequest.Password,
		Name:     updateUserRequest.Name,
		Email:    updateUserRequest.Email,
		Status:   updateUserRequest.Status,
	}

	if err := db.Find(user).Error; err != nil {
		logger.ErrF("Failed to update user: %+v", err.Error())
		return
	}

	updateAccountOptin(updateUserRequest, user)

	if err := db.Save(user).Error; err != nil {
		logger.ErrF("Failed to update user: %+v", err.Error())
		return
	}
	logger.Info("User updated!")
}


func updateAccountOptin(updateUserRequest *model.UpdateUserRequest, user *model.User) {
	updateUserRequestValue := reflect.ValueOf(updateUserRequest).Elem()
	userDomainValue := reflect.ValueOf(user).Elem()

	updateOptinDTOType := updateUserRequestValue.Type()

	for i := 0; i < updateOptinDTOType.NumField(); i++ {
		fieldName := updateOptinDTOType.Field(i).Name
		fieldValue := updateUserRequestValue.Field(i)

		userDomainFieldValue := userDomainValue.FieldByName(fieldName)

		check := !fieldValue.IsZero()
		if check {
			userDomainFieldValue.Set(fieldValue)
		}

	}
}
