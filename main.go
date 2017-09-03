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

	log.Println(transaction)

	body := buildMessage(transaction)
	response, err := twilio.SendMessage(body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintln(w, response.Status)
}

// buildMessage builds the body of the text message to be send
func buildMessage(transaction Transaction) string {
	return fmt.Sprintf("\nType: %s\nDescription: %s\nMerchant: %s\nAmount: %.2f\nCurrency: %s\nCategory: %s\n",
		transaction.Type,
		transaction.description(),
		transaction.merchantName(),
		transaction.amount(),
		transaction.currency(),
		transaction.category())
}
