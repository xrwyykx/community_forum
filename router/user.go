package router

import (
	"communityforum/app/user/controller"

	"github.com/gin-gonic/gin"
)

func setUserRouters(root *gin.RouterGroup) {
	// 用户信息相关
	//root.POST("/common/get-user-info", controller.GetUserInfo)     // 获取用户信息
	//root.POST("/user/update-user-info", controller.UpdateUserInfo) // 更新用户信息
	root.POST("/user/get-user-posts", controller.GetUserPosts) // 获取用户帖子列表

	// 帖子相关
	root.POST("/user/add-post", controller.AddPost)              // 发布帖子
	root.POST("/user/update-post", controller.UpdatePost)        // 修改帖子内容
	root.DELETE("/user/delete-post", controller.DeletePost)      // 删除帖子
	root.POST("/user/get-post-detail", controller.GetPostDetail) // 搜索帖子
	root.POST("/user/reply-post", controller.ReplyPost)          // 回复帖子
	root.POST("/user/get-post-reply", controller.GetPostReply)   // 查看帖子的评论
	root.POST("/user/report-post", controller.ReportPost)        // 举报帖子/评论
	root.POST("/user/report-record", controller.ReportRecord)    // 获取自己的举报记录

	// 关注相关
	root.POST("/user/follow-user", controller.FollowUser)                      // 关注用户
	root.POST("/user/get-follow-or-fans-list", controller.GetFollowOrFansList) // 获取关注列表
}

//[error] failed to parse pending as default value for int, got error: strconv.ParseInt: parsing "pending": invalid syntax
