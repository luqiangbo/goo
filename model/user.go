package model

import (
	"time"
)

type SysUser struct {
	Base
	Name     string    `json:"name" gorm:"comment:姓名"`
	Age      int64     `json:"age" gorm:"comment:年龄"`
	Birthday time.Time `json:"birthday" gorm:"comment:时间"`
}
