package twilio

import (
	"encoding/json"
	"io/ioutil"
)

// Configuration configures the Twillio client
type Config struct {
	APIURL    string
	From      string
	To        string
	BasicAuth struct {
		Username string
		Password string
	}
}

// Init configures the app
func (c *Config) Init() error {
	confData, err := ioutil.ReadFile("twilio/config.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(confData, &c); err != nil {
		return err
	}

	return nil
}
