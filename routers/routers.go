package routers

import (
	"github.com/gin-gonic/gin"
	"tracking/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controllers.IndexHandler)

	return r
}
