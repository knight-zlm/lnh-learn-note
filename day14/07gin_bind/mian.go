package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login ...
type Login struct {
	//  binding:"required"修饰字段是必传字段
	User string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
}

// Pagination ...
type Pagination struct {
	Page     string `form:"page" defualt:"1"`
	PageSize int    `form:"page_size" default:"5"`
}

func main() {
	r := gin.Default()
	r.POST("/loginJson", func(c *gin.Context) {
		var data Login
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		}
		c.JSON(200, gin.H{
			"msg": data,
		})
	})
	r.POST("/loginFrom", func(c *gin.Context) {
		var data Login
		if err := c.Bind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"msg": data,
		})
	})
	r.GET("/articles", func(c *gin.Context) {
		var data Pagination
		if err := c.BindQuery(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"msg": data,
		})
	})
	r.GET("/login/:user", func(c *gin.Context) {
		var data Login
		if err := c.ShouldBindUri(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"msg": data,
		})
	})
	// r.GET("/articles", func(c *gin.Context) {
	// 	var data Login
	// 	if err := c.ShouldBindQuery(&data); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	// 		return
	// 	}
	// 	c.JSON(200, gin.H{
	// 		"msg": data,
	// 	})
	// })
	// r.GET("/someStruct", func(c *gin.Context) {
	// 	var msg struct {
	// 		Name    string
	// 		Message string
	// 		Number  int
	// 	}
	// 	msg.Name = "zhansan"
	// 	msg.Message = "kankan"
	// 	msg.Number = 123
	// 	c.JSON(200, msg)
	// })
	// r.GET("/someXml", func(c *gin.Context) {
	// 	c.XML(200, gin.H{
	// 		"msg": "abc",
	// 	})
	// })
	// r.GET("/someYml", func(c *gin.Context) {
	// 	c.YAML(200, gin.H{
	// 		"name": "zhansan",
	// 	})
	// })
	// r.GET("/someProtubuf", func(c *gin.Context) {
	// 	reps := []int64{1, 2}
	// 	label := "label"
	// 	data := &protoexample.Test{
	// 		Label: &label,
	// 		Reps:  reps,
	// 	}
	// 	c.ProtoBuf(200, data)
	// })
	r.Run(":8080")
}
