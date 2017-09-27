package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	client *Client
	err    error
)

func init() {
	client, err = NewClient()
	if err != nil {
		log.Fatal(err)
	}
}

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

	body := buildMessage(transaction)

	response, err := client.Send(body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintln(w, response.Status)
}

// buildMessage builds the body of the text message to be send
func buildMessage(transaction Transaction) string {
	return transaction.String()
}
