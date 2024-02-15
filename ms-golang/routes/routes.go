package route

import (
	"github.com/caiquefds/ms-users/controller"
	"github.com/caiquefds/ms-users/handler"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	initializeRoutes(router)
	router.Run(":8080")
}

func initializeRoutes(router *gin.Engine) {
	handler.Initialize()

	v1 := router.Group("/users/v1")
	{
		v1.POST("/user", controller.CreateUser)
		v1.GET("/user/:id", controller.GetUserById)
		v1.PATCH("/user/:id", controller.UpdateUser)
		v1.DELETE("/user/:id", controller.DeleteUserById)
	}

}
