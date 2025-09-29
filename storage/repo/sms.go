package repo

import (
	"context"

	"github.com/Hot-One/kizen-go-service/dto"
	"github.com/Hot-One/kizen-go-service/pkg/pg"
)

type SmsInterface interface {
	Create(*dto.CreateSms) (*pg.Id, error)
	Update(*dto.UpdateSms, pg.Filter) error
	FindOne(context.Context, pg.Filter) (*dto.Sms, error)
	Find(context.Context, pg.Filter) ([]dto.Sms, error)
	Page(context.Context, pg.Filter, int64, int64) (*dto.SmsPage, error)
	Delete(pg.Filter) error
}
