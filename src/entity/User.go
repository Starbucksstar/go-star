package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"type:varchar(20);comment:name"`
	Password string `json:"password" form:"password" gorm:"type:varchar(20);comment:密码"`
	City     string `json:"city" gorm:"type:varchar(20);comment:城市"`
}
