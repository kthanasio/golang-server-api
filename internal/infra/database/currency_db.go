package database

import (
	"context"
	"time"

	"github.com/kthanasio/golang-server-api/internal/entity"
	"gorm.io/gorm"
)

var (
	dbMaxTime = 20
)

type Currency struct {
	DB *gorm.DB
}

func NewCurrencyDB(db *gorm.DB) *Currency {
	return &Currency{DB: db}
}

func (c *Currency) Save(ctx context.Context, currency *entity.Currency) error {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(dbMaxTime))
	defer cancel()
	return c.DB.WithContext(ctx).Save(currency).Error
}
