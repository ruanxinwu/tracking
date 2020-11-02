package api

import (
	"github.com/gin-gonic/gin"
	"tracking/utils"
)
import "tracking/model/request"

// 用户登录
func Login(c *gin.Context) {

}

func Register(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(R)

	UserVerify := utils.Rules{
		"Username":    {utils.NotEmpty()},
		"NickName":    {utils.NotEmpty()},
		"Password":    {utils.NotEmpty()},
		"AuthorityId": {utils.NotEmpty()},
	}
	utils.Verify(R, UserVerify)

}
