package controller

import (
	"communityforum/app"
	"communityforum/data/db/user"
	"communityforum/models/co"
	"communityforum/models/dto"
	"github.com/gin-gonic/gin"
	"net/http"
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

func LikesPost(c *gin.Context)    {}
func ReportThings(c *gin.Context) {}
