package controllers

import (
	"github.com/gin-gonic/gin"
	"tracking/global/response"
)

func IndexHandler(c *gin.Context) {
	response.Ok(c)
}
