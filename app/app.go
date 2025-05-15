package app

import (
	"communityforum/global"
	"communityforum/models/dao"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) int64 {
	Session, err := c.Cookie("SESSION")
	if err != nil {
		log.Printf("获取SESSION cookie失败: %v", err)
		return -1
	}
	log.Printf("获取到SESSION cookie: %s", Session)

	redisCil := global.GetRedisConn()
	if redisCil == nil {
		log.Println("Redis客户端为空")
		return -1
	}

	key := global.ProjectName + ":sessions:" + Session
	log.Printf("从Redis获取用户信息，key: %s", key)

	response := redisCil.HGet(c, key, "sessionAttr:user_login")
	if response.Err() != nil {
		log.Printf("从Redis获取用户信息失败: %v", response.Err())
		return -1
	}

	var userSessionValue dao.User
	err = json.Unmarshal([]byte(response.Val()), &userSessionValue)
	if err != nil {
		log.Printf("解析用户信息失败: %v", err)
		return -1
	}

	log.Printf("成功获取用户ID: %d", userSessionValue.UserID)
	return userSessionValue.UserID
}
