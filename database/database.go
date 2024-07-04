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

func InsertCurrencyPair(currencyPair models.CurrencyPair) {
	db.Create(&currencyPair)
}

func UpdateCurrencyPair(currencyPair models.CurrencyPair) {
	db.Save(&currencyPair)
}

func GetCurrencyPair(name string) models.CurrencyPair {
	var currencyPair models.CurrencyPair

	db.First(&currencyPair, "name = ?", name)

	return currencyPair
}

func GetCurrencyPairs() []models.CurrencyPair {
	var currencyPairs []models.CurrencyPair

	db.Find(&currencyPairs)

	return currencyPairs
}

func InsertCurrencyPairMetadata(currencyPairMetadata models.CurrencyPairMetadata) {
	db.Create(&currencyPairMetadata)
}

func UpdateCurrencyPairMetadata(currencyPairMetadata models.CurrencyPairMetadata) {
	db.Save(&currencyPairMetadata)
}

func GetCurrencyPairMetadata(currencyPairID int) models.CurrencyPairMetadata {
	var currencyPairMetadata models.CurrencyPairMetadata

	db.First(&currencyPairMetadata, "currency_pair_id = ?", currencyPairID)

	return currencyPairMetadata
}
