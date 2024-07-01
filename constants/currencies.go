package constants

// Currencies is a list of supported currencies
// [TODO]: a full extensiability thourhg the DB might require to put this list into the DB.
var AvailableCurrencyPairs = []struct {
	PairName    string `json:"pair_name"`
	IsAvailable bool   `json:"is_available"`
}{
	{PairName: "BTC-USD", IsAvailable: true},
	{PairName: "ETH-USD", IsAvailable: true},
	{PairName: "XRP-USD", IsAvailable: true},
}
