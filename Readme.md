# my-monzo

Sends txt messages using [Twilio](https://www.twilio.com)'s API every time you use your [Monzo](https://monzo.com) card for payments.

## Quick how-to

_For deploying to Heroku have a look at the Heroku branch https://github.com/gmichokostas/my-monzo/tree/heroku_

1) First you need to own a [Monzo](https://monzo.com) card.

2) Install [Go](https://golang.org/doc/install) at your machine.

3) Get the source code with `go get github.com/gmichokostas/my-monzo`

4) Create a webhook using [Monzo's API](https://monzo.com/docs/#webhooks) pointing to your URL.

5) Create a [Twilio](https://www.twilio.com) Account.

6) Grab your [Twilio's](https://www.twilio.com) credentials, configure Twilio with your `$EDITOR` and you're ready to go:

    ```
    $ cp twilio/config.json.sample twilio/config.json
    
    $ $EDITOR config.json
    ```

## Configuration

The configuration file contains JSON. 
The following keys can be defined and all are required:

* `APIURL` &ndash; To send a new outgoing message, make an HTTP POST to your Messages list resource URI.
* `From` &ndash; A Twilio phone number for the type of message you wish to send. 
* `To` &ndash; The destination phone number. 
* `Username` &ndash; Basic HTTP Auth username.
* `Password` &ndash; Basic HTTP Auth password.

### Project Purpose

The original motivation for creating this project was firstly to 
be able to get notified every time I use my [Monzo](https://monzo.com) card 
and I don't have access to the internet through my phone and secondly to learn the [Go](https://golang.org/doc/install) language.
