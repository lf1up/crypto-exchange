package models

// User model
type User struct {
	Name string `json:"name"`
}

// [TODO]: write models with GORM usage
type CurrencyPairMetadata struct {
	CurrencyPairID   int    `json:"currency_pair_id"`
	LastInteractedAt string `json:"last_interacted_at"`
}

type CurrencyPair struct {
	Name      string `json:"name"`
	From      string `json:"from"`
	To        string `json:"to"`
	Price     string `json:"price"`
	Market    string `json:"market"`
	UpdatedAt string `json:"date"`
}
