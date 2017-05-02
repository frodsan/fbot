package fbot

import (
	"io"
	"net/http"
)

func Handler(bot Bot) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			validateWebhook(bot, w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func validateWebhook(bot Bot, w http.ResponseWriter, r *http.Request) {
	mode := r.FormValue("hub.mode")
	token := r.FormValue("hub.verify_token")

	if mode == "subscribe" && token == bot.VerifyToken {
		io.WriteString(w, r.FormValue("hub.challenge"))
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}
