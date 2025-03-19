package dao

import "time"

type Notification struct {
	NotifyID   int       `json:"notify_id" gorm:"column:notify_id;autoIncrement"`
	UserID     int       `json:"user_id" gorm:"column:user_id;not null"`
	Type       string    `json:"type" gorm:"column:type;not null"`
	SourceID   int       `json:"source_id" gorm:"column:source_id;not null"`
	IsRead     bool      `json:"is_read" gorm:"column:is_read;default:false"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;not null"`
}

func (a *Notification) TableName() string {
	return "notification"
}
