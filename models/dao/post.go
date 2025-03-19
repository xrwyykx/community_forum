package dao

import "time"

type Post struct {
	PostID     int       `json:"post_id" gorm:"column:post_id;autoIncrement"`
	UserID     int64     `json:"user_id" gorm:"column:user_id;not null"`
	Title      string    `json:"title" gorm:"column:title;not null"`
	Content    string    `json:"content" gorm:"column:content;not null"`
	IsTop      bool      `json:"is_top" gorm:"column:is_top;default:false"`
	IsElite    bool      `json:"is_elite" gorm:"column:is_elite;default:false"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;not null"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

func (a *Post) TableName() string {
	return "post"
}
