package router

import (
	"communityforum/app/comon/controller"
	"github.com/gin-gonic/gin"
)

func setCommonRouters(root *gin.RouterGroup) {
	root.POST("api/common/register", controller.Register)
	root.POST("api/common/login", controller.Login)
	root.POST("api/common/logout", controller.Logout)   //退出登录
	root.DELETE("api/common/logoff", controller.Logoff) //注销账号
	root.POST("api/common/get-user-id", controller.TestGetId)
	//用户信息
	root.POST("api/common/update-info", controller.UpdateUserInfo)
	root.POST("api/common/get-user-info", controller.GetUserInfo)
	root.POST("api/common/test-ai", controller.TestAi)
}
