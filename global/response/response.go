package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time string      `json:"time"`
}

const (
	SUCCESS = 0
	ERROR   = 7
)

func result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
		Time: time.Now().Format("2006-01-02 15:04:05"),
	})
}

func Ok(c *gin.Context) {
	result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	result(ERROR, data, message, c)
}
