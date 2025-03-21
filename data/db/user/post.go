package user

import (
	"communityforum/global"
	"communityforum/models/dao"
	"communityforum/models/dto"
	"github.com/gin-gonic/gin"
)

func GetPostDetail(c *gin.Context, param dto.GetPostMap) (err error, data []dto.PostDetail, total int64) {
	//UserName string `json:"userName" gorm:"column:userName"`
	//Title    string `json:"title" gorm:"column:title"`
	//Content  string `json:"content" gorm:"column:content"`
	db := global.GetDbConn(c).Model(&dao.Post{}).Joins("Join user on post.user_id =user.user_id")
	if param.Content != "" {
		db = db.Where("user.username like ?", "%"+param.Content+"%").
			Where("post.title like ?", "%"+param.Content+"%")
	}
	if err = db.Select("post.*,user.username").Count(&total).Order("post.create_time desc").
		Offset((param.Page.PageIndex - 1) * param.Page.PageSize).
		Limit(param.Page.PageSize).
		Find(&data).Error; err != nil {
		return err, nil, 0
	}
	return
}

//count,order,offset,limit,find
