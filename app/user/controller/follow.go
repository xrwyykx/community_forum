package controller

import (
	"communityforum/app"
	"communityforum/data/db/user"
	"communityforum/models/co"
	"communityforum/models/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FollowUser(c *gin.Context) {
	userId := app.GetUserId(c)
	var param dto.FollowUserMap
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("参数绑定失败"+err.Error()))
		return
	}
	if err := user.Followuser(c, userId, param.FollowedId); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("关注用户失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, co.Success("关注用户成功", nil))
}

func GetFollowOrFansList(c *gin.Context) {
	userId := app.GetUserId(c)
	var param dto.FansOrFollowMap
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("参数绑定失败"+err.Error()))
		return
	}
	err, data, total := user.GetFollowOrFansList(c, userId, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("获取关注/粉丝列表失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, co.Success("获取成功", gin.H{
		"data":  data,
		"total": total,
	}))
}
