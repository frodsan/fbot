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
		fmt.Println(event.Sender.ID)
		fmt.Println(event.Recipient.ID)
		fmt.Println(event.Timestamp)
		fmt.Println(event.Message.Mid)
		fmt.Println(event.Message.Seq)
		fmt.Println(event.Message.Text)

		if len(event.Message.Attachments) != 0 {
			for _, attachment := range event.Message.Attachments {
				fmt.Println(attachment.Type)
				fmt.Println(attachment.Payload.URL)
			}
		}
	})

	fbot.On("delivery", func(event fbot.Event) {
		event.Delivery.Mids[0].Mid // => "mid.1458668856218:ed81099e15d3f4f233"
		event.Delivery.Watermark   // => 1458668856253
		event.Delivery.Seq         // => 37
	})

	fbot.On("postback", func(event fbot.Event) {
		fmt.Println(event.Postback.Payload)
	})

	http.Handle("/bot", fbot.Handler())

	http.ListenAndServe(":4567", nil)
}
