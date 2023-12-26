package main

import (
	"star/src/global"
	"star/src/initialize"
	"star/src/router"
)

func init() {
	println("开始初始化配置")
	initialize.InitConfig()

	println("开始初始化数据库连接")
	initialize.InitGorm()
}

func main() {
	err := router.InitRouter().Run(":9090")
	if err != nil {
		return
	}
	defer closeConnection()
}

func closeConnection() {
	if global.MysqlConfig != nil {
		db, _ := global.GlobalMysqlClient.DB()
		println("开始关闭数据库连接...")
		_ = db.Close()
		println("数据库连接已关闭")
	}
}
