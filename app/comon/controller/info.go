package controller

import (
	"communityforum/app"
	"communityforum/data/db/comon"
	"communityforum/models/co"
	"communityforum/models/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	id := app.GetUserId(c)
	//if err := c.ShouldBindJSON(&param); err != nil {
	//	c.JSON(http.StatusBadRequest, co.BadRequest("参数绑定失败"))
	//	return
	//}
	err, data := comon.GetUserInfo(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("获取用户信息失败"+err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, co.Success("获取用户信息成功", data))
	}
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
