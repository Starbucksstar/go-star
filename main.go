package main

import (
	"star/src/global"
	"star/src/initialize"
	"star/src/router"
	"star/src/scheduler"
	"sync"
)

var once sync.Once

func init() {
	println("开始初始化配置")
	initialize.InitConfig()

	println("开始初始化数据库连接")
	once.Do(initialize.InitGorm)

	println("开始启动定时任务")
	go scheduler.SyncUser()
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
