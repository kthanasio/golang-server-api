package database

import (
	"context"

	"github.com/kthanasio/golang-server-api/internal/entity"
)

type CurrencyInterface interface {
	Save(ctx context.Context, currency *entity.Currency) error
}
