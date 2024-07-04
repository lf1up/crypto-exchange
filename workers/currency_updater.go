package workers

import (
	"crypto-exchange/constants"
	"log"
	"time"

	"github.com/joho/godotenv"
)

var (
	messages = make(chan string)
)

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

	// API_KEY := os.Getenv("FASTFOREX_API_KEY")

	for _, pair := range pairs {
		log.Println("Updating pair: ", pair)

	}
}

func ScheduleBackgroundUpdate(minutes int) {
	time.Sleep(5 * time.Second) // wait a bit before the docker container is fully up and etc.

	for {
		messages <- "all"
		time.Sleep(time.Duration(minutes) * time.Minute)
	}
}
