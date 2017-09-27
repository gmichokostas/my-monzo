package main

import (
	"net/http"

	"github.com/gmichokostas/my-monzo/twilio"
)

type Client struct {
	msgr Messenger
}

func NewClient() (*Client, error) {
	msgr, err := twilio.New()
	if err != nil {
		return nil, err
	}
	return &Client{msgr: msgr}, nil
}

func (c *Client) Send(body string) (*http.Response, error) {
	return c.msgr.Send(body)
}
