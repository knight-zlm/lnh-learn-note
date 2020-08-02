package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行函数
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	// 通配所有的api参数
	r.GET("/all/*name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, name)
	})
	r.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Jack")
		c.String(http.StatusOK, "Hello "+name)
	})
	// 执行
	r.Run(":8000")
}
