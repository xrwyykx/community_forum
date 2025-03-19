package dao

import "time"

type Report struct {
	ReportID   int       `json:"report_id" gorm:"column:report_id;autoIncrement"`
	UserID     int       `json:"user_id" gorm:"column:user_id;not null"`
	TargetID   int       `json:"target_id" gorm:"column:target_id;not null"`
	Type       string    `json:"type" gorm:"column:type;not null"`
	Status     string    `json:"status" gorm:"column:status;default:pending"`
	HandledBy  int       `json:"handled_by" gorm:"column:handled_by"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;not null"`
}

func (a *Report) TableName() string {
	return "report"
}
