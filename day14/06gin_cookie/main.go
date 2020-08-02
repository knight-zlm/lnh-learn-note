package main

import (
	"github.com/gin-gonic/gin"
)

func authCookie(c *gin.Context) {
	// 检查是否有登录cookie
	_, err := c.Cookie("logstatus")
	if err != nil {
		// 没有登录
		c.Set("logstatus", false)
	} else {
		c.Set("logstatus", true)
	}
	c.Next()
	c.Abort()
}

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("logstatus", "true", 60, "/", "localhost", false, true)
		c.JSON(200, "login successed")
	})
	r.GET("/home", authCookie, func(c *gin.Context) {
		value, _ := c.Get("logstatus")
		if value.(bool) {
			c.JSON(200, gin.H{"data": "home"})
		} else {
			c.JSON(200, gin.H{"msg": "statusUnauthorized"})
		}
	})
	// r.GET("/cookie", func(c *gin.Context) {
	// 	// 获取客户端是否携带cookie
	// 	cookie, err := c.Cookie("key_cookie")
	// 	if err != nil {
	// 		cookie = "Notset"
	// 		// maxAge 过期时间 秒为单位
	// 		// path cookie 所在目录
	// 		//domain cookie 所在域名
	// 		//secure 是否只能通过https访问
	// 		//httpOnly 是否允许别人通过ajax获取自己的cookie
	// 		c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
	// 	}
	// fmt.Println("cookie:", cookie)
	// })
	r.Run()
}
