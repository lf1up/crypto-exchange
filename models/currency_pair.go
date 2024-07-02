package models

import (
	"time"

	"gorm.io/gorm"
)

type CurrencyPairMetadata struct {
	gorm.Model
	IsAvailable      bool
	LastInteractedAt time.Time
	CurrencyPairID   int `gorm:"uniqueIndex"`
	CurrencyPair     CurrencyPair
}

type CurrencyPair struct {
	gorm.Model
	Name   string
	From   string
	To     string
	Price  float64
	Market string
}
