package controller

import (
	"net/http"

	"github.com/caiquefds/ms-users/handler"
	"github.com/caiquefds/ms-users/model"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	request := model.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		ValidateRequest(ctx, request, err)
		return
	}

	userId, err := handler.CreateUser(&request)

	if err != nil {
		errorReponse := createErrorMessage(err)
		sendResponseWithBody(ctx, http.StatusUnprocessableEntity, errorReponse)
		return
	}

	url := ctx.Request.Host + ctx.Request.URL.Path + "/" + userId
	ctx.Header("url", url)
	sendResponseWithStatus(ctx, http.StatusNoContent)
}
