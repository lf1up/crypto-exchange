package constants

// EUR, USD, CNY, USDT, USDC, ETH
var AvailableCurrencyPairs = []struct {
	PairName    string `json:"pair_name"`
	IsAvailable bool   `json:"is_available"`
}{
	{PairName: "EUR/USDT", IsAvailable: true},
	{PairName: "EUR/USDC", IsAvailable: true},
	{PairName: "EUR/ETH", IsAvailable: true},
	{PairName: "USD/USDT", IsAvailable: true},
	{PairName: "USD/USDC", IsAvailable: true},
	{PairName: "USD/ETH", IsAvailable: true},
	{PairName: "CNY/USDT", IsAvailable: true},
	{PairName: "CNY/USDC", IsAvailable: true},
	{PairName: "CNY/ETH", IsAvailable: true},
	{PairName: "USDT/EUR", IsAvailable: true},
	{PairName: "USDT/USD", IsAvailable: true},
	{PairName: "USDT/CNY", IsAvailable: true},
	{PairName: "USDC/EUR", IsAvailable: true},
	{PairName: "USDC/USD", IsAvailable: true},
	{PairName: "USDC/CNY", IsAvailable: true},
	{PairName: "ETH/EUR", IsAvailable: true},
	{PairName: "ETH/USD", IsAvailable: true},
	{PairName: "ETH/CNY", IsAvailable: true},
}

func GetAllPairsAsStrings() []string {
	var pairs []string
	for _, pair := range AvailableCurrencyPairs {
		if pair.IsAvailable {
			pairs = append(pairs, pair.PairName)
		}
	}
	return pairs
}
