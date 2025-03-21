package router

import (
	"communityforum/app/comon/controller"
	"github.com/gin-gonic/gin"
)

func setCommonRouters(root *gin.RouterGroup) {
	root.POST("api/common/register", controller.Register)
	root.POST("api/common/login", controller.Login)
	root.POST("api/comon/get-user-id", controller.TestGetId)
	//用户信息
	root.POST("api/comon/update-info", controller.UpdateUserInfo)
	root.POST("api/common/get-user-info", controller.GetUserInfo)
}
