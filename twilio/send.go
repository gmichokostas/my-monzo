package twilio

import (
	"bytes"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// Configuration data for the twillio client
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
	configuration.APIURL = os.Getenv("APIURL")
	configuration.From = os.Getenv("From")
	configuration.To = os.Getenv("To")
	configuration.BasicAuth.Username = os.Getenv("Username")
	configuration.BasicAuth.Password = os.Getenv("Password")
}

// SendMessage sends message to destination number
func SendMessage(body string) (*http.Response, error) {
	request, err := buildRequest(body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return response, nil
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
