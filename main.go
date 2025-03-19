package main

import (
	"communityforum/global"
	"communityforum/router"
)

func main() {
	//加载配置文件
	global.LoadConfig()
	//连接数据库
	global.InitDB()
	//启动路由
	router.InitRouterAndStartServer()
}
