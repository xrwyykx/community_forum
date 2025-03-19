package dao

import "time"

type User struct {
	UserID     int64     `json:"user_id" gorm:"column:user_id;autoIncrement"`
	Username   string    `json:"username" gorm:"column:username;not null"`
	Password   string    `json:"password" gorm:"column:password;not null"`
	Email      string    `json:"email" gorm:"column:email;not null"`
	Avatar     string    `json:"avatar" gorm:"column:avatar;default:default.avatar.jpg"`
	Role       string    `json:"role" gorm:"column:role;default:user"`
	Points     int       `json:"points" gorm:"column:points;default:0"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;not null"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
	LastLogin  time.Time `json:"last_login" gorm:"column:last_login"`
}

func (a *User) TableName() string {
	return "user"
}
