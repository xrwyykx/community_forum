package comon

import (
	"communityforum/global"
	"communityforum/models/dao"
	"communityforum/models/dto"
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

func CheackRegister(c *gin.Context, data dto.RegisterMap) error {
	if data.UserName == "" || data.Password == "" {
		return errors.New("请输入正确的用户名或密码")
	}
	if data.Avatar == "" {
		data.Avatar = "https://th.bing.com/th/id/OIP.B8r7cgNGCrTgXEYhMADqVQAAAA?w=196&h=196&c=7&r=0&o=5&cb=iwc2&dpr=2.2&pid=1.7"
		//return errors.New("请上传一个可爱的头像")
	}
	var count int64
	//查询是否有重复的用户名
	err := global.GetDbConn(c).Model(dao.User{}).Where("username=?", data.UserName).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户名重复咯，你重取一个")
	}

	//对密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//生成id
	userID, err := GenerateSnowflakeID()
	if err != nil {
		return err
	}
	user := dao.User{
		UserID:     userID,
		Username:   data.UserName,
		Password:   string(hashedPassword),
		Email:      data.Email,
		Avatar:     data.Avatar,
		Role:       "user",
		Points:     0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		LastLogin:  time.Now(),
	}
	err = global.GetDbConn(c).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func GenerateSnowflakeID() (int64, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0, err
	}
	return node.Generate().Int64(), nil
}

func CheckLogin(c *gin.Context, data dto.LoginMap) error {
	fmt.Println("接收到的数据", data)
	//userId := app.GetUserId(c)
	if data.UserName == "" || data.Password == "" {
		return errors.New("请输入完整的用户名和密码")
	}
	var user dao.User
	if err := global.GetDbConn(c).Model(&dao.User{}).Where("username = ? ", data.UserName).First(&user).Error; err != nil {
		return errors.New("用户名不存在")
	}
	userId := user.UserID
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return errors.New("用户名或密码错误，请重新输入")
	} else {
		tx := global.GetDbConn(c).Begin()
		param := dao.User{
			LastLogin: time.Now(),
			//Points:    gorm.Expr("points+?", 3),
		}
		err := tx.Model(&dao.User{}).Where("user_id = ?", userId).Updates(&param).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		err = tx.Model(&dao.User{}).Where("user_id = ?", userId).Updates(map[string]interface{}{
			"points": gorm.Expr("points + ?", 3),
		}).Error
		if err != nil {
			return err
		}
		err = tx.Model(&dao.PointLog{}).Where("user_id = ?", userId).Create(map[string]interface{}{
			"user_id":     userId,
			"action":      2,
			"points":      gorm.Expr("points + ?", 3),
			"related_id":  -1,
			"create_time": time.Now(),
		}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.Commit().Error; err != nil {
			return err
		}
		return nil
	}
}

func Logout() {

}
func Logoff(c *gin.Context, userId int64) error {
	var data dao.User
	db := global.GetDbConn(c).Model(&dao.User{})
	if err := db.Where("user_id = ?", userId).First(&data).Error; err != nil {
		return err
	}
	if err := db.Where("user_id = ?", userId).Delete(&dao.User{}).Error; err != nil {
		return err
	}
	return nil
}
