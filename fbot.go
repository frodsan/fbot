package fbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Version is the version of the library.
const Version = "0.0.2"

const (
	// EventMessage represents the message event.
	EventMessage = "message"
	// EventDelivery represents the message-delivery event.
	EventDelivery = "delivery"
	// EventPostback represents the postback event.
	EventPostback = "postback"
)

// Bot is the object that receives and sends
// messages to the Messenger Platform.
type Bot struct {
	Config    *Config
	Callbacks map[string]func(*Event)
}

// Config represents the required configuration to
// receive and send message to the Messenger Platform.
type Config struct {
	AccessToken string
	AppSecret   string
	VerifyToken string
}

// NewBot creates a new instance of Bot.
func NewBot(config Config) *Bot {
	return &Bot{
		Config:    &config,
		Callbacks: make(map[string]func(*Event)),
	}
}

// Event represents the event fired by the webhook.
type Event struct {
	Sender    *User     `json:"sender"`
	Recipient *User     `json:"recipient"`
	Timestamp int64     `json:"timestamp,omitempty"`
	Message   *Message  `json:"message"`
	Delivery  *Delivery `json:"delivery"`
	Postback  *Postback `json:"postback"`
}

// User represents the user that acts like sender or recipient.
type User struct {
	ID          string `json:"id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

// Message represents the message callback object.
type Message struct {
	Mid         string        `json:"mid,omitempty"`
	Seq         int           `json:"seq,omitempty"`
	Text        string        `json:"text,omitempty"`
	Attachment  *Attachment   `json:"attachment,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

// Attachment represents the attachment object included in the message.
type Attachment struct {
	Type    string   `json:"type"`
	Payload *Payload `json:"payload"`
}

// Payload represents the attachment payload data.
type Payload struct {
	URL string `json:"url,omitempty"`
}

// Delivery represents the message-delivered callback object.
type Delivery struct {
	Mids      []string `json:"mids"`
	Watermark int64    `json:"watermark"`
	Seq       int      `json:"seq"`
}

// Postback respresents the postback callback object.
type Postback struct {
	Payload string `json:"payload"`
}

// On registers a callback for the given eventName.
func (bot *Bot) On(eventName string, callback func(*Event)) {
	bot.Callbacks[eventName] = callback
}

func (bot Bot) trigger(event *Event) {
	var eventName string

	if event.Message != nil {
		eventName = EventMessage
	} else if event.Delivery != nil {
		eventName = EventDelivery
	} else if event.Postback != nil {
		eventName = EventPostback
	} else {
		return
	}

	if callback, ok := bot.Callbacks[eventName]; ok {
		callback(event)
	}
}

const baseURL = "https://graph.facebook.com/v2.6/me/messages?access_token=%s"

// DeliverParams represents the message params sent by deliver.
type DeliverParams struct {
	Recipient *User    `json:"recipient"`
	Message   *Message `json:"message"`
}

// Deliver uses the Send API to deliver messages.
func (bot Bot) Deliver(params DeliverParams) error {
	url := fmt.Sprintf(baseURL, bot.Config.AccessToken)

	json, err := json.Marshal(&params)

	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(json))

	if err != nil {
		return err
	}

	return nil
}
