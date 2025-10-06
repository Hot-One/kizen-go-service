package storage

import (
	"github.com/Hot-One/kizen-go-service/pkg/logger"
	"github.com/Hot-One/kizen-go-service/storage/postgres"
	"github.com/Hot-One/kizen-go-service/storage/repo"
	"gorm.io/gorm"
)

type StorageInterface interface {
	Close() error

	Sms() repo.SmsInterface
	User() repo.UserInterface
}

type storage struct {
	db  *gorm.DB
	log logger.Logger

	smsStorage  repo.SmsInterface
	userStorage repo.UserInterface
}

func NewStorage(db *gorm.DB, log logger.Logger) StorageInterface {
	return &storage{
		db:  db,
		log: log,
	}
}

func (s *storage) Close() error {
	pg, err := s.db.DB()
	if err != nil {
		return err
	}

	return pg.Close()
}

func (s *storage) Sms() repo.SmsInterface {
	if s.smsStorage == nil {
		s.smsStorage = postgres.NewSmsStorage(s.db, s.log)
	}

	return s.smsStorage
}

func (s *storage) User() repo.UserInterface {
	if s.userStorage == nil {
		s.userStorage = postgres.NewUserStorage(s.db, s.log)
	}

	return s.userStorage
}
