package db

import "time"

type User struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"first_name" gorm:"type:varchar(255);not null"`
	LastName  string    `json:"last_name" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
