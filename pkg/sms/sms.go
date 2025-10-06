package sms

import (
	"github.com/Hot-One/kizen-go-service/config"
	"github.com/Hot-One/kizen-go-service/pkg/logger"
	"github.com/Hot-One/kizen-go-service/pkg/sms/email"
)

type SmsRepositoryI interface {
	Email() email.EmailI
}

type smsRepository struct {
	cfg *config.Config
	log logger.Logger
}

func NewSmsRepository(cfg *config.Config, log logger.Logger) SmsRepositoryI {
	return &smsRepository{
		cfg: cfg,
		log: log,
	}
}

func (s *smsRepository) Email() email.EmailI {
	return email.NewEmail(s.cfg, s.log)
}
