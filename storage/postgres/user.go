package postgres

import (
	"context"

	"github.com/Hot-One/kizen-go-service/dto"
	models "github.com/Hot-One/kizen-go-service/models/db"
	"github.com/Hot-One/kizen-go-service/pkg/logger"
	"github.com/Hot-One/kizen-go-service/pkg/pg"
	"github.com/Hot-One/kizen-go-service/pkg/security"
	"github.com/Hot-One/kizen-go-service/storage/repo"
	"gorm.io/gorm"
)

type UserStorage struct {
	db  *gorm.DB
	log logger.Logger
}

func NewUserStorage(db *gorm.DB, log logger.Logger) repo.UserInterface {
	return &UserStorage{
		db:  db,
		log: log,
	}
}

func (s *UserStorage) Create(input *dto.CreateUser) (*pg.Id, error) {
	hashedPassword, err := security.HashPassword(input.Password)
	if err != nil {
		s.log.Error("storage: failed to hash password", logger.Error(err))
		return nil, err
	}

	model := &models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hashedPassword,
	}

	if err := pg.Create(s.db, model); err != nil {
		s.log.Error("storage: failed to create user", logger.Error(err))
		return nil, err
	}

	return &pg.Id{Id: model.Id}, nil
}

func (s *UserStorage) Update(input *dto.UpdateUser, filter pg.Filter) error {
	if _, err := pg.Update[models.User](s.db, input, filter); err != nil {
		s.log.Error("storage: failed to update user", logger.Error(err))
		return err
	}

	return nil
}

func (s *UserStorage) FindOne(ctx context.Context, filter pg.Filter) (*dto.User, error) {
	return pg.FindOneWithScan[models.User, dto.User](s.db.WithContext(ctx), filter)
}

func (s *UserStorage) Find(ctx context.Context, filter pg.Filter) ([]dto.User, error) {
	return pg.FindWithScan[models.User, dto.User](s.db.WithContext(ctx), filter)
}

func (s *UserStorage) Page(ctx context.Context, filter pg.Filter, page, size int64) (*dto.UserPage, error) {
	return pg.PageWithScan[models.User, dto.User](s.db.WithContext(ctx), page, size, filter)
}

func (s *UserStorage) Delete(filter pg.Filter) error {
	return pg.Delete[models.User](s.db, nil, filter)
}
