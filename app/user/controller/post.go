package controller

import (
	"communityforum/data/db/user"
	"communityforum/models/co"
	"communityforum/models/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddPost(c *gin.Context) {

}

func UpdatePost(c *gin.Context) {

}
func DeletePost(c *gin.Context) {

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
