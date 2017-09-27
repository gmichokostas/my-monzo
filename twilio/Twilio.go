package twilio

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Twilio client
type Twilio struct {
	config Config
}

func New() (*Twilio, error) {
	c := Config{}

	if err := c.Init(); err != nil {
		return nil, err
	}

	return &Twilio{config: c}, nil
}

func (t Twilio) Send(body string) (response *http.Response, err error) {
	request, err := t.buildRequest(body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	err = retry(3, 2*time.Second, func() error {
		response, err = client.Do(request)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		status := response.StatusCode
		switch {
		case status >= 500:
			return fmt.Errorf("server error: %v", status)
		default:
			return nil
		}
	})

	return
}

// buildRequest builds the request to be send to Twilio
func (t Twilio) buildRequest(body string) (*http.Request, error) {
	data := url.Values{}
	data.Set("To", t.config.To)
	data.Add("From", t.config.From)
	data.Add("Body", body)

	req, err := http.NewRequest("POST", t.config.APIURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(t.config.BasicAuth.Username, t.config.BasicAuth.Password)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	return req, nil
}

// retry to execute the function for attempts number of times
func retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return retry(attempts, 2*sleep, fn)
		}
		return err
	}
	return nil
}
