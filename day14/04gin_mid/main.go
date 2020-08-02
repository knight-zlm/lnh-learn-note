package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// MiddleWare ...
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		c.Set("request", "中间件")
		//执行路由对应的函数
		c.Next()
		// 执行完函数做后续处理
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	r := gin.Default()
	// 加载模板文件
	// r.LoadHTMLGlob("template/*")
	// r.GET("/index", func(c *gin.Context) {
	// 	c.HTML(200, "index.html", gin.H{
	// 		"title": "我的标题",
	// 	})
	// })
	// 注册中间件
	r.Use(MiddleWare())
	// {} 为了代码规范
	{
		r.GET("/middleware", func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			c.JSON(200, gin.H{
				"request": req,
			})
		})
		r.GET("/middleware2", MiddleWare(), func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			c.JSON(200, gin.H{
				"request": req,
			})
		})
	}

	r.Run(":8080")
}
