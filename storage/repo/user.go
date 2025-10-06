package repo

import (
	"context"

	"github.com/Hot-One/kizen-go-service/dto"
	"github.com/Hot-One/kizen-go-service/pkg/pg"
)

type UserInterface interface {
	Create(input *dto.CreateUser) (*pg.Id, error)
	Update(input *dto.UpdateUser, filter pg.Filter) error
	FindOne(ctx context.Context, filter pg.Filter) (*dto.User, error)
	Find(ctx context.Context, filter pg.Filter) ([]dto.User, error)
	Page(ctx context.Context, filter pg.Filter, page, size int64) (*dto.UserPage, error)
	Delete(filter pg.Filter) error
}
