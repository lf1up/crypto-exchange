package database

import (
	"crypto-exchange/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// Connect with database
func Connect(isDev bool) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var host string
	if isDev {
		host = os.Getenv("POSTGRES_HOST")
	} else {
		host = "crypto-api-db"
	}

	dsn := "host=" + host + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect with Database: %v", err)
		return
	}

	db.AutoMigrate(&models.CurrencyPair{})
	db.AutoMigrate(&models.CurrencyPairMetadata{})

	log.Println("Connected with Database!")
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
