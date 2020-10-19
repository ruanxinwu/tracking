package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type responseS interface {
}

type ResponseSuccess struct {
	Code int
	Msg  string
	Now  string
	Data []interface{}
}

type ResponseFail struct {
}

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseSuccess{})
}
