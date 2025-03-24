package router

import (
	"communityforum/app/user/controller"
	"github.com/gin-gonic/gin"
)

func setUserRouters(root *gin.RouterGroup) {
	root.POST("api/user/add-post", controller.AddPost)              //发布帖子
	root.POST("api/user/update-post", controller.UpdatePost)        //修改帖子内容
	root.DELETE("api/user/delete-post", controller.DeletePost)      //删除帖子
	root.POST("api/user/get-post-detail", controller.GetPostDetail) //搜索帖子
}

//Error 1452 (23000): Cannot add or update a child row: a foreign key constraint fails (`community_forum`.`post`, CONSTRAINT `post_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE)"
