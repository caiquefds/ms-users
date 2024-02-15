package controller

import (
	"net/http"

	"github.com/caiquefds/ms-users/handler"
	"github.com/caiquefds/ms-users/model"
	"github.com/gin-gonic/gin"
)

func GetUserById(ctx *gin.Context) {

	request := &model.GetUserByIdRequest{}

	err := ctx.ShouldBindUri(request); if err != nil {
		ValidateRequest(ctx,request,err)
		return
	}

	userId := request.Id

	response, err := handler.GetUserById(userId)

	if err != nil {
		sendResponseWithStatus(ctx, http.StatusBadRequest)
		return
	}

	sendResponseWithBody(ctx, http.StatusOK, response)
}
