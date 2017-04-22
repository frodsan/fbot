package fbot

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Handler returns the handler to use for incoming messages from the
// Facebook Messenger Platform.
func Handler(bot Bot) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			verifyToken(bot, w, r)
		case "POST":
			receiveMessage(bot, w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func verifyToken(bot Bot, w http.ResponseWriter, r *http.Request) {
	if r.FormValue("hub.verify_token") == bot.Config.VerifyToken {
		w.Write([]byte(r.FormValue("hub.challenge")))
	} else {
		w.Write([]byte("Error; wrong verify token"))
	}
}

type receive struct {
	Object  string  `json:"object"`
	Entries []entry `json:"entry"`
}

type entry struct {
	ID     string  `json:"id"`
	Time   int64   `json:"time"`
	Events []Event `json:"messaging"`
}

func receiveMessage(bot Bot, w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Error reading empty response body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	message, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}
	xHubSignature := r.Header.Get("x-hub-signature")

	if xHubSignature == "" || !strings.HasPrefix(xHubSignature, "sha1=") {
		http.Error(w, "Error getting integrity signature", http.StatusBadRequest)
		return
	}

	xHubSignature = xHubSignature[5:] // Remove "sha1=" prefix

	if ok := verifySignature([]byte(xHubSignature), []byte(bot.Config.AppSecret), message); !ok {
		http.Error(w, "Error checking message integrity", http.StatusBadRequest)
		return
	}
	var rec receive

	err = json.Unmarshal(message, &rec)

	if err != nil {
		http.Error(w, "Error parsing response body format", http.StatusBadRequest)
		return
	}

	triggerEvents(bot, rec.Entries)
}

func verifySignature(signature, secret, message []byte) bool {
	mac := hmac.New(sha1.New, secret)
	mac.Write(message)

	expectedSignature := mac.Sum(nil)

	return hmac.Equal(expectedSignature, hexSignature(signature))
}

func hexSignature(signature []byte) []byte {
	s := make([]byte, 20)

	hex.Decode(s, signature)

	return s
}

func triggerEvents(bot Bot, entries []entry) {
	for _, entry := range entries {
		for _, event := range entry.Events {
			bot.trigger(&event)
		}
	}
}
