package router

import (
	"communityforum/app/admin/controller"
	"github.com/gin-gonic/gin"
)

func setAdminRouters(root *gin.RouterGroup) {
	//root.POST("api/admin/get-report-list", controller.GetReportList) //获取举报列表
	root.POST("api/admin/review-report", controller.ReviewRecord) //处理待审核举报
	//这个也还没有完成
}
