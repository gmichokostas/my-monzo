package twilio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Configuration configs twillio client
type Configuration struct {
	APIURL    string
	From      string
	To        string
	BasicAuth struct {
		Username string
		Password string
	}
}

var configuration Configuration

func init() {
	confData, err := ioutil.ReadFile("twilio/config.json")
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(confData, &configuration); err != nil {
		log.Fatalln(err)
	}

}

// SendMessage sends message to destination number
func SendMessage(body string) (response *http.Response, err error) {
	request, err := buildRequest(body)
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
func buildRequest(body string) (*http.Request, error) {
	data := url.Values{}
	data.Set("To", configuration.To)
	data.Add("From", configuration.From)
	data.Add("Body", body)

	req, err := http.NewRequest("POST", configuration.APIURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(configuration.BasicAuth.Username, configuration.BasicAuth.Password)
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
