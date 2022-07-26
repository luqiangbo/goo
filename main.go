package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"goo/middlewares"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
	"time"
)

type User struct {
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Name     string    `json:"name" gorm:"comment:姓名"`
	Age      int64     `json:"age" gorm:"comment:年龄"`
	Birthday time.Time `json:"birthday" gorm:"comment:时间"`
}

func main() {
	// 数据库
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=127.0.0.1 user=postgres password=123456 dbname=ens port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true, //注意：这两行内容是绝对不可以删除的，否则查找不成功
		},
	})

	if err != nil {
		fmt.Println("连接失败")
	}
	M := db.Migrator()

	boolHasUser := M.HasTable(&User{})
	if !boolHasUser {
		errDbInit := M.CreateTable(&User{})
		if errDbInit != nil {
			fmt.Println("创建user表失败")
		}
	}
	r := gin.Default()
	r.Use(middlewares.Cors())
	// 创建
	r.POST("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"type": "post",
		})
	})
	// 删除
	r.DELETE("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"type": "delete",
		})
	})
	// 更新
	r.PUT("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"type": "put",
		})
	})
	// 查询
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"type": "get",
		})
	})

	r.Run(":8001")
	//user := User{UUID: uuid.NewV4(), Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//result := db.Create(&user)
	//fmt.Println(result)
}
