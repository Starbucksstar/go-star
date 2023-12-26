package core

import (
	"fmt"
	"gorm.io/gorm"
	"star/src/entity"
	"star/src/global"
)

func setTableOption(tableComment string) *gorm.DB {
	value := fmt.Sprintf("ENGINE=InnoDB COMMENT='%s'", tableComment)
	return global.GlobalMysqlClient.Set("gorm:table_options", value)
}

// 用户相关表
func userTable() {
	_ = setTableOption("用户表").AutoMigrate(&entity.User{})
}

// 数据表迁移
func AutoMigrate() {
	// 创建用户相关表
	userTable()
}
