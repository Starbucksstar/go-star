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
func migrateTable() {
	_ = setTableOption("用户表").AutoMigrate(&entity.User{})
	_ = setTableOption("角色表").AutoMigrate(&entity.Role{})
}

// 数据表迁移
func AutoMigrate() {
	migrateTable()
}
