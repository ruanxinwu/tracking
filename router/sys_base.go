package router

import (
	"github.com/gin-gonic/gin"
	"tracking/api"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("register", api.Register)
		BaseRouter.POST("login", api.Login)
		BaseRouter.POST("captcha", api.Captcha)
	}
	return BaseRouter
}
