package main

import (
	"fmt"
	"goo/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 跨域
	router.Use(middlewares.Cors())
	router.POST("/upload_chunk", func(c *gin.Context) {
		fileHash := c.PostForm("md5")
		file, err := c.FormFile("file")

		if err != nil {
			fmt.Println("获取文件失败", file, fileHash)
		}
		fmt.Println("获取文件", file, fileHash)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	router.Run()
}
