package database

import (
	"crypto-exchange/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
	// db.AutoMigrate(&models.CurrencyPairMetadata{})

	DB = db

	log.Println("Connected with Database!")
}

func InsertCurrencyPair(currencyPair models.CurrencyPair) {
	if err := DB.Create(&currencyPair).Error; err != nil {
		log.Printf("Error inserting into Database: %v", err)
	}
}

func UpdateCurrencyPair(currencyPair models.CurrencyPair) {
	if err := DB.Save(&currencyPair).Error; err != nil {
		log.Printf("Error updating Database record: %v", err)
	}
}

func GetCurrencyPair(name string) models.CurrencyPair {
	var currencyPair models.CurrencyPair

	result := DB.First(&currencyPair, "name = ?", name)

	if result.Error != nil {
		log.Printf("No Database record found: %v", result.Error)
	}

	return currencyPair
}

func GetCurrencyPairs() []models.CurrencyPair {
	var currencyPairs []models.CurrencyPair

	result := DB.Find(&currencyPairs)

	if result.Error != nil {
		log.Printf("No Database record found: %v", result.Error)
	}

	return currencyPairs
}

// func InsertCurrencyPairMetadata(currencyPairMetadata models.CurrencyPairMetadata) {
// 	DB.Create(&currencyPairMetadata)
// }

// func UpdateCurrencyPairMetadata(currencyPairMetadata models.CurrencyPairMetadata) {
// 	DB.Save(&currencyPairMetadata)
// }

// func GetCurrencyPairMetadata(currencyPairID int) models.CurrencyPairMetadata {
// 	var currencyPairMetadata models.CurrencyPairMetadata

// 	DB.First(&currencyPairMetadata, "currency_pair_id = ?", currencyPairID)

// 	return currencyPairMetadata
// }
