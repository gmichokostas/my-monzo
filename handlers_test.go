package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_transactionHandler(t *testing.T) {

	req, err := http.NewRequest("POST", "/transaction", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transactionHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
