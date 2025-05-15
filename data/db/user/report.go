package user

import (
	"communityforum/global"
	"communityforum/models/dao"
	"communityforum/models/dto"
	"communityforum/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

func ReportPost(c *gin.Context, userId int64, param dto.ReportThingsMap) (err error) {
	if param.Kinds == 0 {
		var res dao.Post
		if err := global.GetDbConn(c).Model(&dao.Post{}).Where("post_id = ?", param.TargetId).First(&res).Error; err != nil {
			return err
		}
	} else {
		var res dao.Comment
		if err := global.GetDbConn(c).Model(&dao.Comment{}).Where("comment_id = ?", param.TargetId).First(&res).Error; err != nil {
			return err
		}
	}
	//检查一下是否举报过了，检查一下巨报表里面是不是有一样的
	var total int64
	if err := global.GetDbConn(c).Model(&dao.Report{}).Where("user_id = ? and target_id = ? and kinds = ?", userId, param.TargetId, param.Kinds).Count(&total).Error; err != nil {
		return nil
	}
	if total != 0 {
		return errors.New("已经举报了，请稍后")
	}
	data := dao.Report{
		UserID:     userId,
		TargetID:   param.TargetId,
		Type:       param.Type,
		Status:     0,
		HandledBy:  global.AdminId,
		CreateTime: time.Now(),
		Kinds:      param.Kinds,
	}
	if err := global.GetDbConn(c).Model(&dao.Report{}).Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetReportRecord(c *gin.Context, userId int64) (data []dto.ReportRecordList, err error, total int64) {
	err = global.GetDbConn(c).Model(&dao.Report{}).Where("user_id = ?", userId).Select("report.target_id,report.kinds,report.create_time,report.status").Count(&total).Find(&data).Error
	if err != nil {
		return nil, err, 0
	}
	for i := range data {
		var content string
		if data[i].Kinds == 0 {
			if err := global.GetDbConn(c).Model(&dao.Post{}).Where("post_id = ?", data[i].TargetID).Select("post.title as content").First(&content).Error; err != nil {
				return nil, err, 0
			}
		} else {
			if err := global.GetDbConn(c).Model(&dao.Comment{}).Where("comment_id = ?", data[i].TargetID).Select("comment.content").First(&content).Error; err != nil {
				return nil, err, 0
			}
		}
		data[i].Content = content
		data[i].CreateTimeMar = utils.MarshalTime(data[i].CreateTime)
	}
	return data, nil, total
}
