package main

import (
	"communityforum/global"
	"communityforum/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//加载配置文件
	global.LoadConfig()
	gin.SetMode(gin.ReleaseMode)
	//连接redis
	global.InitRedis()
	//连接数据库
	global.InitDB()
	//启动路由
	router.InitRouterAndStartServer()
}
