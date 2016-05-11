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
	fbot.Configure(fbot.Config{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		AppSecret:   os.Getenv("APP_SECRET"),
		VerifyToken: os.Getenv("VERIFY_TOKEN"),
	})

	fbot.On("message", func(event fbot.Event) {
		fmt.Println(event.Message.Text)
	})

	http.Handle("/bot", fbot.Handler())

	http.ListenAndServe(":4567", nil)
}
```

API
---

### fbot.Configure(c Config)

Configures the bot with the application's access token,
app secret, and verify token.

```go
fbot.Configure(fbot.Config{
	AccessToken: os.Getenv("ACCESS_TOKEN"),
	AppSecret:   os.Getenv("APP_SECRET"),
	VerifyToken: os.Getenv("VERIFY_TOKEN"),
})
```

## fbot.Handler()

Returns the `http.Handler` that receives the request sent by the Messenger platform.

```go
http.Handle("/bot", fbot.Handler())
```

## fbot.On(eventName string, callback func(Event))

Registers a `callback` for the given `eventName`.

```go
fbot.On("message", func(event fbot.Event) {
	event.Sender.ID    // => 1234567890
	event.Recipient.ID // => 0987654321
	event.Timestamp    // => 1462966178037

	event.Message.Mid  // => "mid.1234567890:41d102a3e1ae206a38"
	event.Message.Seq  // => 41
	event.Message.Text // => "Hello World!"

	event.Message.Attachments[0].Type        // => "image"
	event.Message.Attachments[0].Payload.URL // => https://scontent.xx.fbcdn.net/v/t34.0-12/...
})
```

Panics if `eventName` is not one of the following: `"message"`, `"delivery"`, `"postback"`, or `"optin"`.

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
