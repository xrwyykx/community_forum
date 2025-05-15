package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouterAndStartServer() {
	router := gin.Default()
	// 使用自定义CORS中间件
	router.Use(CorsHandler())
	//路由组是gin中的一个概念，允许将有相关的路由组织在一起，为他们设置相同的中间件或前缀路径
	root := router.Group(viper.GetString("http.path"))
	//使用viper的函数来获取配置文件中的path参数（config.toml）
	setCommonRouters(root) //注册一些通用的组件
	setAdminRouters(root)
	setUserRouters(root)
	router.Run(":8090")
}
func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许的源
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, OPTIONS, DELETE")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Expose-Headers", "Set-Cookie, Content-Length, Content-Range")

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// 打印请求信息用于调试
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL)
		log.Printf("Origin: %s", c.Request.Header.Get("Origin"))
		log.Printf("Cookie: %s", c.Request.Header.Get("Cookie"))
		log.Printf("Path: %s", c.Request.URL.Path)

		c.Next()
	}
}
