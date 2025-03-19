package dao

import "time"

type PointLog struct {
	Action     string    `json:"action"`
	CreateTime time.Time `json:"createTime"`
	LogId      int       `json:"logId"`
	Points     int       `json:"points"`
	RelatedId  int       `json:"RelatedId"`
	UserId     int64     `json:"UserId"`
}

func (a *PointLog) TableName() string {
	return "point_log"
}
