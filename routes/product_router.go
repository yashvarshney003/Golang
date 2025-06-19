package routes

import (
	"golang_crud/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.Engine) {
	productGroup := r.Group("/products")
	{
		productGroup.GET("/", controllers.GetProducts)
		productGroup.POST("/", controllers.InsertProducts)
	}

}
