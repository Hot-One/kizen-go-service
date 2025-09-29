package models

import (
	modelsDB "github.com/Hot-One/kizen-go-service/models/db"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		modelsDB.Sms{},
	)
}
