package controller

import (
	"communityforum/data/db/comon"
	"communityforum/models/co"
	"communityforum/models/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data dto.LoginMap
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("参数绑定失败"+err.Error()))
		return
	}
	if err := comon.CheckLogin(c, data); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("用户登录失败"+err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, co.Success("登录成功", nil))
	}
}
func Register(c *gin.Context) {
	var data dto.RegisterMap //获取到用户信息然后根据信息来进行注册
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("参数绑定失败"+err.Error()))
		return
	}
	if err := comon.CheackRegister(c, data); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("用户注册失败"+err.Error()))
		return

	} else {
		c.JSON(http.StatusOK, co.Success("用户注册成功", nil))
	}
}
