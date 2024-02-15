package controller

import (
	"github.com/gin-gonic/gin"
)

func sendResponseWithBody(ctx *gin.Context, code int, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, data)
}

func sendResponseWithStatus(ctx *gin.Context, code int) {
	ctx.Header("Content-type", "application/json")
	ctx.Status(code)
}
