package main

import (
	"github.com/gin-gonic/gin"
	"goo/middlewares"
)

func main() {
	router := gin.Default()
	// 跨域
	router.Use(middlewares.Cors())
	// 测试
	router.GET("/test", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "ok",
		})

	})

	router.Run()
}
