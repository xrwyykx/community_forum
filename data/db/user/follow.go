package user

import (
	"communityforum/global"
	"communityforum/models/dao"
	"communityforum/models/dto"
	"communityforum/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func Followuser(c *gin.Context, userId int64, followedId int64) (err error) {
	var data dao.User
	if err := global.GetDbConn(c).Model(&dao.User{}).Where("user_id = ?", followedId).First(&data).Error; err != nil {
		return err
	}
	res := dao.Follow{
		FollowedID: followedId,
		FollowerID: userId,
		CreateTime: time.Now(),
	}
	if err := global.GetDbConn(c).Model(&dao.Follow{}).Create(&res).Error; err != nil {
		return err
	}
	return nil
}

func GetFollowOrFansList(c *gin.Context, userId int64, param dto.FansOrFollowMap) (err error, data []dto.GetFansOrFollowList, total int64) {
	db := global.GetDbConn(c).Model(&dao.Follow{})
	//kind 0 fans 1 follow
	if param.Kinds == 0 { //fans 所以当前是 关注当前用户的 当前用户作为followed_id
		db = db.Joins("join user on user.user_id = follow.follower_id").Where("follow.followed_id = ?", userId)
	} else {
		db = db.Joins("join user on user.user_id = follow.followed_id").Where("follow.follower_id = ?", userId)
	}
	if param.UserId != 0 {
		db = db.Where("user_id = ?", param.UserId) //被关注者id
	}
	if param.UserName != "" {
		db = db.Where("username like ?", "%"+param.UserName+"%")
	}
	if err := db.Select("user.username,user.user_id,user.last_login,user.avatar,follow.*").Count(&total).Order("follow.create_time desc").Find(&data).Error; err != nil {
		return err, nil, 0
	}
	for i := range data {
		data[i].LastLoginMar = utils.MarshalTime(data[i].LastLogin)
	}
	return
}
