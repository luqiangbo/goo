package main

import (
	"fmt"
	"goo/middlewares"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	//isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}

func main() {
	router := gin.Default()
	// 跨域
	router.Use(middlewares.Cors())
	// 分片上传
	router.POST("/upload_chunk", func(c *gin.Context) {
		fileHash := c.Request.FormValue("file_hash")
		index := c.Request.FormValue("index")
		total := c.Request.FormValue("total")
		file, errFile := c.FormFile("file")
		if errFile != nil {
			c.JSON(200, gin.H{
				"message": "error 1",
			})
			return
		}

		pathHash := fmt.Sprintf("upload/%s", fileHash)
		os.MkdirAll(pathHash, os.ModePerm)
		errSave := c.SaveUploadedFile(file, fmt.Sprintf("upload/%s/%s", fileHash, index))
		if errSave != nil {
			c.JSON(200, gin.H{
				"message": "error 3",
			})
			return
		}

		c.JSON(200, gin.H{
			"message":  "ok",
			"index":    index,
			"total":    total,
			"file_ash": fileHash,
		})
	})

	// 分片合并
	router.GET("/merge_chunk", func(c *gin.Context) {
		hash := c.Query("hash")
		fileName := c.Query("file_name")
		hashPath := fmt.Sprintf("upload/%s", hash)
		isExistPath, err := PathExists(hashPath)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "hash路径错误",
			})
			return
		}
		if !isExistPath {
			c.JSON(200, gin.H{
				"message":     "文件夹不存在",
				"isExistPath": isExistPath,
				"hashPath":    hashPath,
			})
			return
		}
		isExistPathFile, errFile := PathExists(hashPath + "/" + fileName)
		if errFile != nil {
			c.JSON(200, gin.H{
				"message": "文件名错误",
			})
			return
		}
		if isExistPathFile {
			c.JSON(200, gin.H{
				"message": "文件已存在",
			})
			return
		}

		fileList, errList := ioutil.ReadDir(hashPath)
		if errList != nil {
			c.JSON(200, gin.H{
				"message": "合并文件读取失败",
			})
			return
		}

		// 创建文件
		complateFile, err := os.Create(hashPath + "/" + fileName)
		// 记得关闭文件
		defer complateFile.Close()

		for _, u := range fileList {
			fileBuffer, err := ioutil.ReadFile(hashPath + "/" + u.Name())
			if err != nil {
				fmt.Println("文件打开错误")
			}
			complateFile.Write(fileBuffer)
		}

		c.JSON(200, gin.H{
			"message":     "ok",
			"hash":        hash,
			"fileName":    fileName,
			"isExistPath": isExistPath,
		})
	})

	router.Run()
}
