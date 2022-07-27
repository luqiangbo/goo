package initialize

import (
	"fmt"
	"goo/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Gorm() *gorm.DB {
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
		fmt.Println("db连接失败")
		return nil
	}
	M := db.Migrator()
	boolHasUser := M.HasTable(&model.SysUser{})
	if !boolHasUser {
		errDbInit := M.CreateTable(&model.SysUser{})
		if errDbInit != nil {
			fmt.Println("创建user表失败")
		}
	}
	return db
}
