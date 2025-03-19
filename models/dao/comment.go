package dao

import "time"

type Comment struct {
	CommentID  int       `json:"comment_id" gorm:"column:comment_id;autoIncrement"`
	UserID     int64     `json:"user_id" gorm:"column:user_id;not null"`
	PostID     int       `json:"post_id" gorm:"column:post_id;not null"`
	ParentID   int       `json:"parent_id" gorm:"column:parent_id;default:null"`
	Content    string    `json:"content" gorm:"column:content;not null"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;not null"`
}

func (a *Comment) TableName() string {
	return "comment"
}
