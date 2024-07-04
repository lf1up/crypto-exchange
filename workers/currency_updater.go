package workers

import (
	"crypto-exchange/constants"
	"fmt"
	"time"
)

var (
	messages = make(chan string)
)

func StartCurrencyUpdater() {
	fmt.Println("Currency Updating Worker has Started!")

	for {
		select {
		case msg, ok := <-messages:
			if !ok {
				fmt.Println("Channel closed, exiting.")
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
	fmt.Println("Updating currency pairs...")

	for _, pair := range pairs {
		// [TODO]: update the currency pair in the database from https://www.fastforex.io/ API
		fmt.Println("Updating pair: ", pair)
	}
}

func ScheduleBackgroundUpdate(minutes int) {
	time.Sleep(5 * time.Second) // wait a bit before the docker container is fully up and etc.

	for {
		messages <- "all"
		time.Sleep(time.Duration(minutes) * time.Minute)
	}
}
