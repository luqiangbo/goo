package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goo/global"
	"goo/initialize"
	"goo/middlewares"
	"goo/model"
	"net/http"
	"time"
)

func main() {
	// 数据库

	global.GVA_DB = initialize.Gorm()

	if global.GVA_DB == nil {
		fmt.Println("db初始化失败")
		return
	}

	r := gin.Default()
	r.Use(middlewares.Cors())
	// 创建
	r.POST("/user", func(c *gin.Context) {
		var req model.SysUser
		_ = c.ShouldBindJSON(&req)

		user := model.SysUser{Name: req.Name, Age: req.Age, Birthday: time.Now()}
		err := global.GVA_DB.Create(&user).Error
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"type": "post",
			"req":  req,
			"err":  err,
		})
	})
	// 删除
	r.DELETE("/user", func(c *gin.Context) {
		var req model.SysUser
		_ = c.ShouldBindJSON(&req)
		err := global.GVA_DB.Delete(&model.SysUser{}, req.ID).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"type":    "delete",
				"message": "删除失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"type": "delete",
		})
	})
	// 更新
	r.PUT("/user", func(c *gin.Context) {
		var req model.SysUser
		_ = c.ShouldBindJSON(&req)
		global.GVA_DB.Model(&model.SysUser{}).Where("id = ?", req.ID).Update("age", req.Age)
		c.JSON(http.StatusOK, gin.H{
			"type": "put",
		})
	})
	// 查询
	r.GET("/user", func(c *gin.Context) {
		var list []model.SysUser
		err := global.GVA_DB.Find(&list).Error
		c.JSON(http.StatusOK, gin.H{
			"type": "get",
			"err":  err,
			"list": list,
		})
	})

	r.Run(":8001")

}
