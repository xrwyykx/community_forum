package router

import (
	"communityforum/app/comon/controller"
	"github.com/gin-gonic/gin"
)

func setCommonRouters(root *gin.RouterGroup) {
	root.POST("api/common/get-user-info", controller.GetUserInfo)
}
