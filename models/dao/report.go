package dao

import "time"

type Report struct {
	ReportID   int       `json:"report_id" gorm:"column:report_id;autoIncrement"`
	UserID     int64     `json:"user_id" gorm:"column:user_id;not null"`
	TargetID   int       `json:"target_id" gorm:"column:target_id;not null"`
	Type       int       `json:"type" gorm:"column:type;not null"`
	Status     int       `json:"status" gorm:"column:status;"`
	HandledBy  int64     `json:"handled_by" gorm:"column:handled_by"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;not null"`
	Kinds      int       `json:"kinds" gorm:"column:kinds"`
	ReportedId int64     `json:"reported_id" gorm:"reported_id"` //被举报者id
}

func (a *Report) TableName() string {
	return "report"
}
