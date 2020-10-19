package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type responseS interface {
}

type ResponseSuccess struct {
	Code int
	msg  string
	now  string
	data []interface{}
}

type ResponseFail struct {
}

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseSuccess{})
}
