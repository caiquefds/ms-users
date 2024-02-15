package controller

import (
	"net/http"

	"github.com/caiquefds/ms-users/handler"
	"github.com/caiquefds/ms-users/model"
	"github.com/gin-gonic/gin"
)

func UpdateUser(ctx *gin.Context) {
	request := model.UpdateUserRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ValidateRequest(ctx, request, err)
		return
	}

	handler.UpdateUser(&request)
	sendResponseWithStatus(ctx, http.StatusNoContent)
}
