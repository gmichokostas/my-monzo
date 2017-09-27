package main

import "net/http"

// Messenger is responsible for sending txt messages
type Messenger interface {
	Send(body string) (response *http.Response, err error)
}
