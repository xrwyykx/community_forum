package controller

import (
	"communityforum/app"
	"communityforum/data/db/comon"
	"communityforum/models/co"
	"communityforum/models/dto"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	id := app.GetUserId(c)
	if id <= 0 {
		c.JSON(http.StatusUnauthorized, co.BadRequest("未登录或会话已过期"))
		return
	}

	err, data := comon.GetUserInfo(c, id)
	if err != nil {
		log.Printf("获取用户信息失败: user_id=%d, error=%v", id, err)
		c.JSON(http.StatusBadRequest, co.BadRequest("获取用户信息失败: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, co.Success("获取用户信息成功", data))
}

func UpdateUserInfo(c *gin.Context) {
	var param dto.UpdateInfoMap
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("参数绑定失败"+err.Error()))
		return
	}
	err := comon.UpdateUserIfo(c, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("修改用户信息失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, co.Success("修改用户信息成功", nil))
}
