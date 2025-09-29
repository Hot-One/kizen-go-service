package postgres

import (
	"context"

	"github.com/Hot-One/kizen-go-service/dto"
	models "github.com/Hot-One/kizen-go-service/models/db"
	"github.com/Hot-One/kizen-go-service/pkg/logger"
	"github.com/Hot-One/kizen-go-service/pkg/pg"
	"github.com/Hot-One/kizen-go-service/storage/repo"
	"gorm.io/gorm"
)

type SmsStorage struct {
	db  *gorm.DB
	log logger.Logger
}

func NewSmsStorage(db *gorm.DB, log logger.Logger) repo.SmsInterface {
	return &SmsStorage{
		db:  db,
		log: log,
	}
}

func (s *SmsStorage) Create(input *dto.CreateSms) (*pg.Id, error) {
	model := &models.Sms{
		Type:  input.Type,
		Value: input.Value,
		Code:  input.Code,
	}

	if err := pg.Create(s.db, model); err != nil {
		s.log.Error("storage: failed to create sms", logger.Error(err))
		return nil, err
	}

	return &pg.Id{Id: model.Id}, nil
}

func (s *SmsStorage) Update(input *dto.UpdateSms, filter pg.Filter) error {
	if _, err := pg.Update[models.Sms](s.db, input, filter); err != nil {
		s.log.Error("storage: failed to update sms", logger.Error(err))
		return err
	}

	return nil
}

func (s *SmsStorage) FindOne(ctx context.Context, filter pg.Filter) (*dto.Sms, error) {
	return pg.FindOneWithScan[models.Sms, dto.Sms](s.db.WithContext(ctx), filter)
}

func (s *SmsStorage) Find(ctx context.Context, filter pg.Filter) ([]dto.Sms, error) {
	return pg.FindWithScan[models.Sms, dto.Sms](s.db.WithContext(ctx), filter)
}

func (s *SmsStorage) Page(ctx context.Context, filter pg.Filter, page, size int64) (*dto.SmsPage, error) {
	return pg.PageWithScan[models.Sms, dto.Sms](s.db.WithContext(ctx), page, size, filter)
}

func (s *SmsStorage) Delete(filter pg.Filter) error {
	return pg.Delete[models.Sms](s.db, nil, filter)
}
