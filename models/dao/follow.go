package dao

import "time"

type Follow struct {
	FollowID   int       `json:"follow_id" gorm:"column:follow_id;autoIncrement"`
	FollowerID int64     `json:"follower_id" gorm:"column:follower_id;not null"`
	FollowedID int64     `json:"followed_id" gorm:"column:followed_id;not null"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;not null"`
}

func (a *Follow) TableName() string {
	return "follow"
}
