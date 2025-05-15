package controller

import (
	"github.com/gin-gonic/gin"
)

//func GetReportList(c *gin.Context) {
//	userId := app.GetUserId(c)
//	var param dto.GetReportListMap
//	if err := c.ShouldBindJSON(&param); err != nil {
//		c.JSON(http.StatusBadRequest, co.BadRequest("参数绑定失败"+err.Error()))
//		return
//	}
//	data, err, total := admin.GetReportList(c, userId, param)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, co.BadRequest("获取举报记录失败"+err.Error()))
//		return
//	}
//	c.JSON(http.StatusOK, co.Success("获取举报记录成功", gin.H{
//		"data":  data,
//		"total": total,
//	}))
//}

func ReviewRecord(c *gin.Context) {

}
