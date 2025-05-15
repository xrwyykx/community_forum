package controller

import (
	"communityforum/app"
	"communityforum/data/db/comon"
	"communityforum/global"
	"communityforum/models/co"
	"communityforum/models/dao"
	"communityforum/models/dto"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		session := generateSession()
		saveSessionToRedis(c, session, data.UserName)

		// 设置cookie，确保在所有路径下可访问
		c.SetCookie("SESSION", session, 3600, "/", "", false, false)

		// 打印调试信息
		log.Printf("Setting cookie: SESSION=%s", session)

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

func generateSession() string {
	return uuid.New().String()
}

func saveSessionToRedis(c *gin.Context, session string, userName string) {
	redisClient := global.GetRedisConn()
	var user dao.User
	if err := global.GetDbConn(c).Model(&dao.User{}).Where("username = ?", userName).First(&user).Error; err != nil {
		log.Println("获取用户信息失败")
		return
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		log.Println("序列化用户信息失败:", err)
		return
	}
	//err = redisClient.HSet(c, global.ProjectName+"sessions:"+session, "sessionAttr:user_login", string(userJson)).Err()
	err = redisClient.HSet(c, global.ProjectName+":sessions:"+session, "sessionAttr:user_login", string(userJson)).Err()
	if err != nil {
		log.Println("将sessioin存入redis失败:", err)
	}
}

func Logout(c *gin.Context) {

}

func Logoff(c *gin.Context) {
	userId := app.GetUserId(c)
	if err := comon.Logoff(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, co.BadRequest("注销账号失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, co.Success("注销账号成功", nil))
}
