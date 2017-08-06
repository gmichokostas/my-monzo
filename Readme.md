# my-monzo

Sends txt messages using [Twilio](https://www.twilio.com)'s API every time you use your [Monzo](https://monzo.com) card for payments.

## Quick how-to

This branch provides information about deploying to [Heroku](https://heroku.com).

1) First you need to own a [Monzo](https://monzo.com) card.

2) Install [Go](https://golang.org/doc/install) at your machine.

3) Get the source code with `go get github.com/gmichokostas/my-monzo`

4) Create a webhook using [Monzo's API](https://monzo.com/docs/#webhooks) pointing to your URL.

5) Create a [Twilio](https://www.twilio.com) Account.

## Configuration

Grab your [Twilio's](https://www.twilio.com) credentials and set the following config vars:

* `APIURL` &ndash; To send a new outgoing message, make an HTTP POST to your Messages list resource URI.
* `From` &ndash; A Twilio phone number for the type of message you wish to send. 
* `To` &ndash; The destination phone number. 
* `Username` &ndash; Basic HTTP Auth username.
* `Password` &ndash; Basic HTTP Auth password.

You can read this [guide](https://devcenter.heroku.com/articles/config-vars) for how to set up config vars 
with Heroku CLI.

### Project Purpose

The original motivation for creating this project was firstly to 
be able to get notified every time I use my [Monzo](https://monzo.com) card 
and I don't have access to the internet through my phone and secondly to learn the [Go](https://golang.org/doc/install) language.
