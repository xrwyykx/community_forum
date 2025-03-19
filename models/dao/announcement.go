package dao

import "time"

type Announcement struct {
	AnnounceId int       `json:"announceId" gorm:"column:announceId"`
	Content    string    `json:"content" gorm:"column:content"`
	CreateTime time.Time `json:"createTime" gorm:"column:"createTime`
	Title      string    `json:"title" gorm:"column:title"`
}

func (a *Announcement) TableName() string {
	return "announcement"
}
