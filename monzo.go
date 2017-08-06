package main

// Transaction created from Monzo
type Transaction struct {
	Type string
	Data Data
}

// Data of the transaction
type Data struct {
	Amount      int
	Currency    string
	Description string
	Category    string
	Merchant    Merchant
}

// Merchant info
type Merchant struct {
	Name string
}
