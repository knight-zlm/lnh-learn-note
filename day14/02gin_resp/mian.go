package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
	r := gin.Default()
	r.GET("/someJson", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":    "someJson",
			"status": 200,
		})
	})
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "zhansan"
		msg.Message = "kankan"
		msg.Number = 123
		c.JSON(200, msg)
	})
	r.GET("/someXml", func(c *gin.Context) {
		c.XML(200, gin.H{
			"msg": "abc",
		})
	})
	r.GET("/someYml", func(c *gin.Context) {
		c.YAML(200, gin.H{
			"name": "zhansan",
		})
	})
	r.GET("/someProtubuf", func(c *gin.Context) {
		reps := []int64{1, 2}
		label := "label"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})
	r.Run(":8080")
}
