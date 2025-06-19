package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	UserRouter(r)
	ProductRouter(r)
	return r

}
