package storage

import (
	"github.com/Hot-One/kizen-go-service/pkg/logger"
	"gorm.io/gorm"
)

type StorageInterface interface {
}

type storage struct {
	db  *gorm.DB
	log logger.Logger
}

func NewStorage(db *gorm.DB, log logger.Logger) StorageInterface {
	return &storage{
		db:  db,
		log: log,
	}
}
