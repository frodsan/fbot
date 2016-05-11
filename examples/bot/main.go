package main

import (
	"fmt"
	"net/http"
	"os"

	"../../../fbot"
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
