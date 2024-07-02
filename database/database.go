package database

import (
	"crypto-exchange/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// Connect with database
func Connect() {
	// [TODO]: put credential into an .env file
	dsn := "host=crypto-api-db user=postgres password=lolwut123 dbname=exchange port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// [TODO]: handle error
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.CurrencyPair{})
	db.AutoMigrate(&models.CurrencyPairMetadata{})

	fmt.Println("Connected with Database")
}

func Insert(currencyPair models.CurrencyPair) {
	db.Create(&currencyPair)
}

func Get() []models.CurrencyPair {
	var currencyPairs []models.CurrencyPair

	db.Find(&currencyPairs)

	return currencyPairs
}
