package fbot

import "fmt"

// Event represents the event fired by the webhook.
type Event struct {
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	Timestamp int64     `json:"timestamp,omitempty"`
	Message   *Message  `json:"message"`
}

// Sender represents the user who sent the message.
type Sender struct {
	ID int64 `json:"id"`
}

// Recipient represents the user who the message was sent to.
type Recipient struct {
	ID int64 `json:"id"`
}

// Message represents the message callback object.
type Message struct {
	Mid         string       `json:"mid"`
	Seq         int          `json:"seq"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

// Attachment represents the attachment object included in the message.
type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

// Payload represents the attachment payload data.
type Payload struct {
	URL string `json:"url,omitempty"`
}

// Delivery represents the message-delivered callback object.
type Delivery struct {
	Mids      []string `json:"mids"`
	WaterMark int64    `json:"watermark"`
	Seq       int      `json:"seq"`
}

// Postback respresents the postback callback object.
type Postback struct {
	Payload string `json:"payload"`
}

var callbacks = make(map[string]func(Event))

// On registers a hook for the given event.
func On(eventName string, callback func(Event)) {
	if isPermitted(eventName) {
		callbacks[eventName] = callback
	} else {
		panic(fmt.Sprintf("'%s' is not a valid event", eventName))
	}
}

var permittedEvents = []string{"message", "delivery", "postback", "optin"}

func isPermitted(eventName string) bool {
	for _, e := range permittedEvents {
		if e == eventName {
			return true
		}
	}

	return false
}

func trigger(key string, event Event) {
	if callback, ok := callbacks[key]; ok {
		callback(event)
	}
}
