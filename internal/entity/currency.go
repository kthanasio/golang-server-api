package entity

import "gorm.io/gorm"

type Currency struct {
	gorm.Model
	Date  string
	Value float64
}

func NewCurrency(date string, value float64) *Currency {
	return &Currency{Date: date, Value: value}
}
