package workers

import (
	"crypto-exchange/constants"
	"crypto-exchange/database"
	"crypto-exchange/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var (
	messages = make(chan string)
)

type APIDataResponse struct {
	IsError bool
	Rate    float64
	From    string
	To      string
}

func StartCurrencyUpdater() {
	log.Println("Currency Updating Worker has Started!")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	for {
		select {
		case msg, ok := <-messages:
			if !ok {
				log.Println("Channel closed, exiting.")
				return
			}
			if msg == "all" {
				go UpdateCurrencyPairs(constants.GetAllPairsAsStrings())
			} else {
				go UpdateCurrencyPairs([]string{msg})
			}
		case <-time.After(2 * time.Second):
			// no signal received
			continue
		}
	}
}

// Use this function to signal the currency updater to do an immediate update.
func SignalCurrencyUpdater(msg string) {
	messages <- msg
}

func UpdateCurrencyPairs(pairs []string) {
	log.Println("Updating currency pairs...")

	for _, pair := range pairs {
		log.Println("Updating pair:", pair)

		result := FetchAPIData(pair)
		if !result.IsError {
			pairInstance := database.GetCurrencyPair(pair)
			if pairInstance.ID == 0 {
				database.InsertCurrencyPair(models.CurrencyPair{
					Name: pair,
					From: result.From,
					To:   result.To,
					Rate: result.Rate,
				})
			} else {
				pairInstance.Rate = result.Rate
				database.UpdateCurrencyPair(pairInstance)
			}
		}

		log.Println("Pair updated!")
	}

	log.Println("Currency pairs have been updated!")
}

func FetchAPIData(pair string) APIDataResponse {
	from := strings.Split(pair, "/")[0]
	to := strings.Split(pair, "/")[1]
	apiKey := os.Getenv("FASTFOREX_API_KEY")
	url := fmt.Sprintf("https://api.fastforex.io/convert?from=%s&to=%s&amount=1.00&api_key=%s", from, to, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("API fetching error: %v", err)
		return APIDataResponse{IsError: true, Rate: 0, From: from, To: to}
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("API response decoding error: %v", err)
		return APIDataResponse{IsError: true, Rate: 0, From: from, To: to}
	}

	if resultData, ok := result["result"].(map[string]interface{}); ok {
		if rateValue, ok := resultData["rate"].(float64); ok {
			return APIDataResponse{IsError: false, Rate: rateValue, From: from, To: to}
		}
	}

	return APIDataResponse{IsError: true, Rate: 0, From: from, To: to}
}

func ScheduleBackgroundUpdate(minutes int) {
	time.Sleep(5 * time.Second) // wait a bit before the docker container is fully up and etc.

	for {
		messages <- "all"
		time.Sleep(time.Duration(minutes) * time.Minute)
	}
}
