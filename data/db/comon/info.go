package comon

import (
	"communityforum/global"
	"communityforum/models/dao"
	"communityforum/models/dto"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUserInfo(c *gin.Context, id int64) (err error, data dto.GetInfoMap) {
	err = global.GetDbConn(c).Model(&dao.User{}).Where("user_id = ?", id).First(&data).Error
	if err != nil {
		return err, dto.GetInfoMap{}
	}
	return nil, data
}

func UpdateUserIfo(c *gin.Context, param dto.UpdateInfoMap) error {
	var user dao.User
	if err := global.GetDbConn(c).Model(&dao.User{}).Where("username = ?", param.UserName).First(&user).Error; err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if err := global.GetDbConn(c).Model(&dao.User{}).Where("user_id = ?", param.UserId).Updates(map[string]interface{}{
		"username": param.UserName,
		"password": string(hashedPassword),
		"email":    param.Email,
		"avatar":   param.Avatar,
	}).Error; err != nil {
		return err
	}
	return nil
}
