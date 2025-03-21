package controller

import (
	"communityforum/app"
	"communityforum/models/co"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestGetId(c *gin.Context) {
	id := app.GetUserId(c)
	if id == -1 {
		c.JSON(http.StatusBadRequest, co.BadRequest("获取用户id失败"))
		//c.JSON(http.StatusOK, co.Success("获取用户id失败", -1))
		return
	}
	c.JSON(http.StatusOK, co.Success("成功获取用户id", id))
}
