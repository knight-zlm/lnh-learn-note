package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func taskTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	since := time.Since(start)
	fmt.Printf("url:%s task:%v\n", c.Request.URL, since)
}

func main() {
	r := gin.Default()
	// 创建组
	staRouter := r.Group("/staTime", taskTime)
	{
		staRouter.GET("/index", staIndexHandler)
		staRouter.GET("/home", staHomeHandler)
	}
	r.GET("/outindex", outIndexHandler)
	r.Run()
}

func staIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func staHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}

func outIndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "out"})
}
