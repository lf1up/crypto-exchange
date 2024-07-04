package models

import (
	"gorm.io/gorm"
)

// [WARNING]: Unused at this moment, but could become useful in the very near future.
// type CurrencyPairMetadata struct {
// 	gorm.Model
// 	IsAvailable      bool
// 	LastInteractedAt time.Time
// 	CurrencyPairID   int `gorm:"uniqueIndex"`
// 	CurrencyPair     CurrencyPair
// }

type CurrencyPair struct {
	gorm.Model
	Name string
	From string
	To   string
	Rate float64
}
