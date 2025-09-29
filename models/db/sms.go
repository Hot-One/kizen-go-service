package db

import "time"

type Sms struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	Type      string    `json:"type" gorm:"type:varchar(255);not null"`
	Value     string    `json:"value" gorm:"type:varchar(255);not null"`
	Code      string    `json:"code" gorm:"type:varchar(10);not null"`
	SentCount int       `json:"sent_count" gorm:"default:0"`
	SentAt    time.Time `json:"sent_at"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
