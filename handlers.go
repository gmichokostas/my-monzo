package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	client, err := NewClient()
	if err != nil {
		http.Error(w, "could not create client: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var transaction Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "could not decode request: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	body := transaction.String()
	response, err := client.Send(body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintln(w, response.Status)
}
