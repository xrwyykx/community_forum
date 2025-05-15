package controller

import (
	"communityforum/app"
	"communityforum/data/db/user"
	"communityforum/models/co"
	"communityforum/models/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReportPost(c *gin.Context) {
	userId := app.GetUserId(c)
	var param dto.ReportThingsMap
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("参数绑定失败"+err.Error()))
		return
	}
	err := user.ReportPost(c, userId, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("举报失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, co.Success("举报成功", nil))
}

// 获取自己的举报记录
func ReportRecord(c *gin.Context) {
	userId := app.GetUserId(c)
	data, err, total := user.GetReportRecord(c, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("获取举报记录失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, co.Success("获取举报记录成功", gin.H{
		"data":  data,
		"total": total,
	}))
}
