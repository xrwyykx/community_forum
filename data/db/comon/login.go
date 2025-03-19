package comon

import (
	"communityforum/global"
	"communityforum/models/dao"
	"communityforum/models/dto"
	"errors"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func CheackRegister(c *gin.Context, data dto.RegisterMap) error {
	//有些约束条件
	//1必须输入用户名/密码/头像  如果没有邮箱也没事
	//if err := c.ShouldBindJSON(&data); err != nil {
	//	return c.JSON(http.StatusOK, gin.H{
	//		"message":"参数绑定错误",
	//	})//怎么返回错误
	//}
	//用户名不能重复
	//密码不能小于6位
	//要对密码进行加密
	//2 如果前面都顺利通过，需要生成一个user_id,通过雪花算法生成，避免用户id复用
	//如果以上都可以，将变化加载到数据库，user表
	if data.UserName == "" || data.Password == "" {
		return errors.New("请输入正确的用户名或密码")
	}
	if data.Avatar == "" {
		return errors.New("请上传一个可爱的头像")
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
	//var count int64
	if data.UserName == "" || data.Password == "" {
		return errors.New("请输入完整的用户名和密码")
	}
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	return err
	//}
	//println(string(hashedPassword))
	var user dao.User
	if err := global.GetDbConn(c).Model(&dao.User{}).Where("username = ? ", data.UserName).First(&user).Error; err != nil {
		return errors.New("用户名不存在")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return errors.New("用户名或密码错误，请重新输入")
	} else {
		param := dao.User{
			LastLogin: time.Now(),
		}
		err := global.GetDbConn(c).Model(&dao.User{}).Where("username = ?", data.UserName).Updates(&param).Error
		if err != nil {
			return err
		}
		return nil
	}
}
