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

func (t Transaction) description() string {
	return t.Data.Description
}

func (t Transaction) amount() float64 {
	return float64(t.Data.Amount) / float64(100)
}

func (t Transaction) currency() string {
	return t.Data.Currency
}
func (t Transaction) category() string {
	return t.Data.Category
}

func (t Transaction) merchantName() string {
	return t.Data.Merchant.Name
}
