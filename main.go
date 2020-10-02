package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "jack")
		c.String(200, fmt.Sprintf("hello: %s", name))
	})
	r.Run(":8083")
}
