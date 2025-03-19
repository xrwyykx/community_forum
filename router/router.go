package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouterAndStartServer() {
	router := gin.Default()
	//路由组是gin中的一个概念，允许将有相关的路由组织在一起，为他们设置相同的中间件或前缀路径
	root := router.Group(viper.GetString("http.path"))
	//使用viper的函数来获取配置文件中的path参数（config.toml）
	setCommonRouters(root) //注册一些通用的组件
	setAdminRouters(root)
	setUserRouters(root)
}
