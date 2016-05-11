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

Design
------

The API is heavily inspired by [hyperoslo/facebook-messenger].

License
-------

fbot is released under the [MIT License].

[Messenger Platform]: https://developers.facebook.com/docs/messenger-platform
[MIT License]: http://opensource.org/licenses/MIT
[hyperoslo/facebook-messenger]: https://github.com/hyperoslo/facebook-messenger
