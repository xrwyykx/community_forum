package controller

import (
	"communityforum/app"
	"communityforum/data/db/user"
	"communityforum/models/co"
	"communityforum/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPost(c *gin.Context) {
	//还有除了发布帖子帖子表会发生变化，要注意用户表的用户积分
	//也会发生变化，因为用户可以通过发帖子增加积分
	userId := app.GetUserId(c)
	var param dto.AddPostMap
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("参数绑定失败"+err.Error()))
		return
	}
	if err := user.AddPost(c, userId, param); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("用户发布帖子失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, co.Success("用户成功发布帖子", nil))
}

func UpdatePost(c *gin.Context) {
	//userId := app.GetUserId(c)
	var param dto.UpdatePostMap
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("参数绑定失败"))
		return
	}
	if err := user.UpdatePost(c, param); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("编辑帖子失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, co.Success("更新帖子成功", nil))
}
func DeletePost(c *gin.Context) {
	var param dto.PostIdMap
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("参数绑定失败"+err.Error()))
		return
	}
	err := user.DeletePost(c, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("删除帖子失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, co.Success("删除帖子成功", nil))
}
func GetPostDetail(c *gin.Context) {
	var param dto.GetPostMap
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("参数绑定失败"+err.Error()))
		return
	}
	err, data, total := user.GetPostDetail(c, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("获取帖子信息失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, co.Success("获取帖子信息成功", gin.H{
		"data":  data,
		"total": total,
	}))
}
func ReplyPost(c *gin.Context) {
	userId := app.GetUserId(c)
	var param dto.ReplyPostMap
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("参数绑定失败"+err.Error()))
		return
	}
	if err := user.ReplyPost(c, param, userId); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("回复帖子失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, co.Success("回复帖子成功", nil))
}

func GetPostReply(c *gin.Context) {
	var param dto.PostIdMap
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("参数绑定失败"+err.Error()))
		return
	}
	data, err, total := user.GetPostReply(c, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("获取帖子评论失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, co.Success("获取帖子评论成功", gin.H{
		"data":  data,
		"total": total,
	}))
}

func LikesPost(c *gin.Context) {
	//var param dto.
}

func GetUserPosts(c *gin.Context) {
	//userId := app.GetUserId(c)
	//if userId <= 0 {
	//	c.JSON(http.StatusUnauthorized, co.BadRequest("未登录或会话已过期"))
	//	return
	//}
	//
	//var param dto.PageParam
	//if err := c.ShouldBindJSON(&param); err != nil {
	//	c.JSON(http.StatusBadRequest, co.BadRequest("参数绑定失败: "+err.Error()))
	//	return
	//}
	//
	//posts, total, err := comon.GetUserPosts(c, userId, param.Page, param.PageSize)
	//if err != nil {
	//	log.Printf("获取用户帖子失败: user_id=%d, error=%v", userId, err)
	//	c.JSON(http.StatusBadRequest, co.BadRequest("获取用户帖子失败: "+err.Error()))
	//	return
	//}
	//
	//c.JSON(http.StatusOK, co.Success("获取用户帖子成功", gin.H{
	//	"total": total,
	//	"posts": posts,
	//}))
}
