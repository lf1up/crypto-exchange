package workers

import (
	"crypto-exchange/constants"
)

func CurrencyUpdater() {
	// Connected with database
	for _, pair := range constants.AvailableCurrencyPairs {
		// [TODO]: spawn a go routine here to update the currency pair in the database
		// [TIP]: you can use time.Sleep() to set the time interval for updating the currency pair
	}
}
