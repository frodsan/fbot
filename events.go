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

// Message represents a Messenger text message.
type Message struct {
	Mid  string `json:"mid"`
	Seq  int    `json:"seq"`
	Text string `json:"text"`
}

var hooks = make(map[string]func(Event))

// On registers a hook for the given event.
func On(eventName string, fn func(Event)) {
	if isPermitted(eventName) {
		hooks[eventName] = fn
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
	if fn, ok := hooks[key]; ok {
		fn(event)
	}
}
