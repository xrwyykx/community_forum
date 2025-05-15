package dto

import "time"

type FollowUserMap struct {
	FollowedId int64 `json:"followed_id" gorm:"column:followedId"` //被关注者id
}

type FansOrFollowMap struct {
	Kinds    int    `json:"kinds" gorm:"column:kinds"` //0是fans 1follow 默认是fans
	UserName string `json:"username" gorm:"column:username"`
	UserId   int64  `json:"user_id" gorm:"column:userId"`
}

type GetFansOrFollowList struct {
	Avatar       string    `json:"avatar" gorm:"column:avatar"`
	UserName     string    `json:"username" gorm:"column:username"`
	UserId       int64     `json:"user_id" gorm:"column:user_id"`
	LastLogin    time.Time `json:"-" gorm:"column:last_login"`
	LastLoginMar string    `json:"last_login"`
}
