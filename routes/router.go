package routes

import (
	"golang_crud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.POST("/", controllers.InsertUsers)
	}
	return r

}
