package controller

import (
	"net/http"

	"github.com/caiquefds/ms-users/handler"
	"github.com/caiquefds/ms-users/model"
	"github.com/gin-gonic/gin"
)

func DeleteUserById(ctx *gin.Context) {

	request := &model.DeleteUserByIdRequest{}

	err := ctx.ShouldBindUri(request); if err != nil {
		ValidateRequest(ctx,request,err)
		return
	}


	userId := request.Id

	handler.DeleteUserById(userId)

	sendResponseWithStatus(ctx, http.StatusNoContent)
}
