package router

import (
	"communityforum/app/comon/controller"
	"github.com/gin-gonic/gin"
)

func setCommonRouters(root *gin.RouterGroup) {
	root.POST("api/common/get-user-info", controller.GetUserInfo)
	//agentOuterService(root.Group("api/common"))
	root.POST("api/common/register", controller.Register)
	root.POST("api/common/login", controller.Login)
}

//func agentOuterService(apiPath *gin.RouterGroup){
//	apiPath.POST("api/common/register", controller.Register)
//
//}
