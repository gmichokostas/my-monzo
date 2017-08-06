package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gmichokostas/my-monzo/twilio"
)

func main() {
	http.HandleFunc("/transaction", transaction)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func transaction(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	body := message(&transaction)
	response, err := twilio.SendMessage(body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintln(w, response.Status)
}

// message builds the body of the text message to be send
func message(transaction *Transaction) string {
	return fmt.Sprintf("\nType: %s\nDescription: %s\nMerchant: %s\nAmount: %d\nCurrency: %s\nCategory: %s\n",
		transaction.Type,
		transaction.Data.Description,
		transaction.Data.Merchant.Name,
		transaction.Data.Amount,
		transaction.Data.Currency,
		transaction.Data.Category)
}
