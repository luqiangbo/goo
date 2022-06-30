package main

import (
	"fmt"
	"goo/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 跨域
	router.Use(middlewares.Cors())
	router.POST("/upload_chunk", func(c *gin.Context) {
		file_hash := c.Request.FormValue("file_hash")
		index := c.Request.FormValue("index")
		total := c.Request.FormValue("total")
		file, errFile := c.FormFile("file")
		if errFile != nil {
			c.JSON(200, gin.H{
				"message": "error 1",
			})
			return
		}
		path_hash := fmt.Sprintf("./upload/%s", file_hash)
		os.Mkdir(path_hash, os.ModePerm)
		errSave := c.SaveUploadedFile(file, fmt.Sprintf("./upload/%s/%s", file_hash, file.Filename))
		if errSave != nil {
			c.JSON(200, gin.H{
				"message": "error 2",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "ok",
			"index":   index,
			"total":   total,
			"hash1":   file_hash,
		})
	})
	router.Run()
}
