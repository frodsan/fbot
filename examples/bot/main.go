package main

import (
	"fmt"
	"net/http"
	"os"

	"../../../fbot"
)

func main() {
	bot := fbot.NewBot(fbot.Config{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		AppSecret:   os.Getenv("APP_SECRET"),
		VerifyToken: os.Getenv("VERIFY_TOKEN"),
	})

	bot.On(fbot.EventMessage, func(event *fbot.Event) {
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

	bot.On(fbot.EventDelivery, func(event *fbot.Event) {
		fmt.Println(event.Delivery.Mids[0])
		fmt.Println(event.Delivery.Watermark)
		fmt.Println(event.Delivery.Seq)
	})

	bot.On(fbot.EventPostback, func(event *fbot.Event) {
		fmt.Println(event.Postback.Payload)
	})

	http.Handle("/bot", fbot.Handler(bot))

	http.ListenAndServe(":4567", nil)
}
