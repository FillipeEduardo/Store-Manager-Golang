package models

import "time"

type Sale struct {
	Id   int32     `gorm:"primaryKey;autoIncrement"`
	Date time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
}
