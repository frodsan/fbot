fbot
====

Bots for Facebook Messenger.

Description
-----------

A simple library for making bots for the [Messenger Platform].

Installation
------------

```
$ go get github.com/frodsan/fbot
```

Usage
-----

```go
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/frodsan/fbot"
)

func main() {
	bot := fbot.NewBot(fbot.Config{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		AppSecret:   os.Getenv("APP_SECRET"),
		VerifyToken: os.Getenv("VERIFY_TOKEN"),
	})

	bot.On(fbot.EventMessage, func(event *fbot.Event) {
		fmt.Println(event.Message.Text)
	})

	http.Handle("/bot", fbot.Handler(bot))

	http.ListenAndServe(":4567", nil)
}
```

API
---

### fbot.NewBot(c Config)

NewBot creates a new instance of a bot with the application's access token,
app secret, and verify token.

```go
bot := fbot.NewBot(fbot.Config{
	AccessToken: os.Getenv("ACCESS_TOKEN"),
	AppSecret:   os.Getenv("APP_SECRET"),
	VerifyToken: os.Getenv("VERIFY_TOKEN"),
})
```

## fbot.Handler(bot *fb.Bot)

Returns the `http.Handler` that receives the request sent by the Messenger platform.

```go
http.Handle("/bot", fbot.Handler(bot))
```

## (\*fb.Bot) On(eventName string, callback func(\*Event))

Registers a `callback` for the given `eventName`.

```go
bot.On(fbot.EventMessage, func(event *fbot.Event) {
	event.Sender.ID    // => 1234567890
	event.Recipient.ID // => 0987654321
	event.Timestamp    // => 1462966178037

	event.Message.Mid  // => "mid.1234567890:41d102a3e1ae206a38"
	event.Message.Seq  // => 41
	event.Message.Text // => "Hello World!"

	event.Message.Attachments[0].Type        // => "image"
	event.Message.Attachments[0].Payload.URL // => https://scontent.xx.fbcdn.net/v/t34.0-12/...
})

bot.On(fbot.EventDelivery, func(event *fbot.Event) {
	event.Delivery.Mids[0]   // => "mid.1458668856218:ed81099e15d3f4f233"
	event.Delivery.Watermark // => 1458668856253
	event.Delivery.Seq       // => 37
})

bot.On(fbot.EventPostback, func(event *fbot.Event) {
	event.Postback.Payload // => "{foo:'foo',bar:'bar'}"
})
```

Configuration
-------------

Follow the [Messenger Platform quickstart] guide for set up the needed Facebook page and development app.

Development
-----------

To test the bot locally, use [ngrok].

Design
------

The API is heavily inspired by [hyperoslo/facebook-messenger].

License
-------

fbot is released under the [MIT License].

[hyperoslo/facebook-messenger]: https://github.com/hyperoslo/facebook-messenger
[Messenger Platform]: https://developers.facebook.com/docs/messenger-platform
[Messenger Platform quickstart]: https://developers.facebook.com/docs/messenger-platform/quickstart
[MIT License]: http://opensource.org/licenses/MIT
[ngrok]: https://ngrok.com/
