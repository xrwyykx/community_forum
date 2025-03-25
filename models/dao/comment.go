package dao

import "time"

type Comment struct {
	CommentID  int       `json:"comment_id" gorm:"column:comment_id;autoIncrement"`
	UserID     int64     `json:"user_id" gorm:"column:user_id;not null"`         //评论者id
	PostID     int       `json:"post_id" gorm:"column:post_id;not null"`         //所属帖子id
	ParentID   int       `json:"parent_id" gorm:"column:parent_id;default:null"` //如果是0说明回复的是帖子，否则这个指的是被回复的评论的id
	Content    string    `json:"content" gorm:"column:content;not null"`         //回复内容
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;not null"` //回复时间
}

func (a *Comment) TableName() string {
	return "comment"
}
