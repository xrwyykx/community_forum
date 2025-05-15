package dto

import (
	"time"
)

// 用户
type ReportThingsMap struct {
	TargetId int `json:"target_id" gorm:"column:targetId"` //被举报内容id
	Type     int `json:"type" gorm:"column:type" `         //内容类型（0涉嫌敏感内容 1传播谣言 2人身攻击 3垃圾广告
	Kinds    int `json:"kinds" gorm:"column:kinds"`        //举报的内容类型（0帖子 1评论）
}
type ReportRecordList struct {
	//UserName      string    `json:"username" gorm:"column:userName"` //被举报的人的名
	TargetID      int       `json:"target_id" gorm:"column:target_id;not null"`
	Kinds         int       `json:"kinds" gorm:"column:kinds"`
	Content       string    `json:"content" gorm:"column:content"` //被举报内容，根据kinds和target_id来进行推断
	CreateTime    time.Time `json:"-" gorm:"column:create_time"`
	CreateTimeMar string    `json:"create_time"`
	Status        int       `json:"status" gorm:"column:status"`
}

// 管理员
type GetReportListMap struct {
	Status int `json:"status" gorm:"column:status"`
}

type ReportList struct {
	UserId int64 `json:"user_id" gorm:"column:user_id"`
	//TargetId      int       `json:"target_id" gorm:"column:target_id"`
	//Type          int       `json:"type" gorm:"column:type"`
	Status        int       `json:"status" gorm:"column:status"`
	CreateTime    time.Time `json:"-" gorm:"column:create_time"`
	CreateTimeMar string    `json:"create_time"`
	//Kinds         int       `json:"kinds" gorm:"column:kinds"`
}

type ReportDetails struct {
	Content       string    `json:"content" gorm:"column:content"`
	UserId        int64     `json:"user_id" gorm:"column:user_id"`
	TargetId      int       `json:"target_id" gorm:"column:target_id"`
	Type          int       `json:"type" gorm:"column:type"`
	Status        int       `json:"status" gorm:"column:status"`
	CreateTime    time.Time `json:"-" gorm:"column:create_time"`
	CreateTimeMar string    `json:"create_time"`
	Kinds         int       `json:"kinds" gorm:"column:kinds"`
}

type ReviewReport struct {
	ReportId   int   `json:"report_id" gorm:"column:report_id"`
	Status     int   `json:"status" gorm:"column:status"`
	UserId     int64 `json:"user_id" gorm:"column:user_id"`  //举报者加分
	ReportedId int64 `json:"reported_id" gorm:"reported_id"` //被举报者扣分，如果举报到为<0,那么就实施封号，不允许发帖，还有就是修改每次登录都会加分
}
