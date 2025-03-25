package user

import (
	"communityforum/global"
	"communityforum/models/dao"
	"communityforum/models/dto"
	"communityforum/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func GetPostDetail(c *gin.Context, param dto.GetPostMap) (err error, data []dto.PostDetail, total int64) {
	//UserName string `json:"userName" gorm:"column:userName"`
	//Title    string `json:"title" gorm:"column:title"`
	//Content  string `json:"content" gorm:"column:content"`
	db := global.GetDbConn(c).Model(&dao.Post{}).Joins("Join user on post.user_id =user.user_id")
	if param.Content != "" {
		if param.Status == 0 { //不填默认为0 也就是按标题或者内容
			db = db.Where("post.title like ? OR post.content like ?", "%"+param.Content+"%", "%"+param.Content+"%")
		} else if param.Status == 1 {
			db = db.Where("user.username like ?", "%"+param.Content+"%")
		}
	}
	if err = db.Select("post.*,user.username,user.user_id").Count(&total).Error; err != nil {
		return err, nil, 0
	}
	if err = db.Select("post.*, user.username, user.user_id").
		//Count(&total).
		Order("post.create_time desc").
		Offset((param.Page.PageIndex - 1) * param.Page.PageSize).
		Limit(param.Page.PageSize).
		Find(&data).Error; err != nil {
		return err, nil, 0
	}
	for i := range data {
		data[i].CreateTimeMar = utils.MarshalTime(data[i].CreateTime)
	}
	return
}

func AddPost(c *gin.Context, userId int64, param dto.AddPostMap) (err error) {
	if param.Content == "" {
		return errors.New("文章内容不允许为空")
	} else if param.Title == "" {
		return errors.New("请正确输入文章标题")
	}
	//解决两个操作需要同时成功的情况，如果一个不成功就会回滚之前的操作
	//开启事务
	tx := global.GetDbConn(c).Begin()
	if tx.Error != nil {
		return tx.Error
	}
	//创建帖子
	data := dao.Post{Title: param.Title,
		UserID:     userId,
		Content:    param.Content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err = tx.Model(&dao.Post{}).Create(&data).Error; err != nil {
		tx.Rollback() //如果创建失败，回滚帖子
		return err
	}
	//更新用户积分
	if err = tx.Model(&dao.User{}).Where("user_id = ?", userId).Updates(map[string]interface{}{
		"points": gorm.Expr("points + ?", 5),
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	//更新积分记录表
	res := dao.PointLog{Points: 5,
		RelatedId:  data.PostID,
		Action:     0,
		UserId:     userId,
		CreateTime: time.Now(),
	}
	if err = tx.Model(&dao.PointLog{}).Where("user_id = ?", userId).Create(&res).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

// 解决如果传来的更新值为空，会将原来的值赋值为空的问题
func UpdatePost(c *gin.Context, param dto.UpdatePostMap) error {
	var data dao.Post
	db := global.GetDbConn(c).Model(&dao.Post{}).Where("post_id=?", param.PostId)
	if err := db.First(&data).Error; err != nil {
		return err
	}
	updateField := map[string]interface{}{}
	if param.Content != "" {
		updateField["content"] = param.Content
	}
	if param.Title != "" {
		updateField["title"] = param.Title
	}
	if len(updateField) > 0 {
		if err := db.Updates(updateField).Error; err != nil {
			return err
		}
	}
	return nil
}

func DeletePost(c *gin.Context, param dto.PostIdMap) error {
	var data dao.Post
	if err := global.GetDbConn(c).Model(&dao.Post{}).Where("post_id = ?", param.PostId).First(&data).Error; err != nil {
		return err
	}
	if err := global.GetDbConn(c).Model(&dao.Post{}).Where("post_id = ?", param.PostId).Delete(&dao.Post{}).Error; err != nil {
		return err
	}
	return nil
}

func ReplyPost(c *gin.Context, param dto.ReplyPostMap, userId int64) error {
	var data dao.Post
	if err := global.GetDbConn(c).Model(&dao.Post{}).Where("post_id = ?", param.PostId).First(&data).Error; err != nil {
		return err
	}
	tx := global.GetDbConn(c).Begin()
	if err := tx.Error; err != nil {
		return err
	}
	res := dao.Comment{UserID: userId, PostID: param.PostId, ParentID: param.ParentId, Content: param.Content, CreateTime: time.Now()}
	if err := tx.Model(&dao.Comment{}).Create(&res).Error; err != nil {
		tx.Rollback()
		return err
	}
	//Updates(map[string]interface{}{
	//	"points": gorm.Expr("points + ?", 5),
	if err := tx.Model(&dao.User{}).Where("user_id = ?", userId).Updates(map[string]interface{}{
		"points": gorm.Expr("points+?", 3),
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	res2 := dao.PointLog{UserId: userId, Action: 1, Points: 3, RelatedId: param.PostId, CreateTime: time.Now()}
	if err := tx.Model(&dao.PointLog{}).Create(&res2).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func GetPostReply(c *gin.Context, param dto.PostIdMap) (data []dto.PostReply, err error, total int64) {
	var res dao.Post
	if err = global.GetDbConn(c).Model(&dao.Post{}).Where("post_id = ?", param.PostId).First(&res).Error; err != nil {
		return
	}
	err = global.GetDbConn(c).Model(&dao.Comment{}).Where("post_id = ?", param.PostId).
		Joins("join user on user.user_id = comment.user_id").
		Select("comment.*,user.user_id,user.username").
		Count(&total).
		Order("comment.create_time desc").
		Find(&data).Error
	if err != nil {
		return
	}

	for i := range data {
		data[i].CreateTimeMar = utils.MarshalTime(data[i].CreateTime)
	}
	return
}

//count,order,offset,limit,find
