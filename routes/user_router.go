package routes

import (
	"golang_crud/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.POST("/", controllers.InsertUsers)
	}

}
