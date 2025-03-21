package dto

import (
	"communityforum/models"
	"time"
)

type GetPostMap struct {
	models.Page
	Content string `json:"content" gorm:"column:content"`
}

type PostDetail struct {
	UserName   string    `json:"userName" gorm:"column:userName"`
	Title      string    `json:"title" gorm:"column:title;not null"`
	Content    string    `json:"content" gorm:"column:content;not null"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;not null"`
	IsTop      bool      `json:"is_top" gorm:"column:is_top;default:false"`
	IsElite    bool      `json:"is_elite" gorm:"column:is_elite;default:false"`
}
