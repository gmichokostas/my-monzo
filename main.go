package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/transaction", transactionHandler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
