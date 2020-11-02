package initialize

import (
	"github.com/gin-gonic/gin"
	"tracking/middleware"
	"tracking/router"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	Router.Use(middleware.Cors())

	ApiGroup := Router.Group("")
	router.InitBaseRouter(ApiGroup)
	return Router
}
