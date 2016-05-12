package fbot

import "fmt"

// Version is the version of the library.
const Version = "0.0.1"

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

// On registers a callback for the given eventName.
func (bot *Bot) On(eventName string, callback func(*Event)) {
	if isValidEvent(eventName) {
		bot.Callbacks[eventName] = callback
	} else {
		panic(fmt.Sprintf("'%s' is not a valid event", eventName))
	}
}

var validEvents = []string{"message", "delivery", "postback"}

func isValidEvent(eventName string) bool {
	for _, e := range validEvents {
		if e == eventName {
			return true
		}
	}

	return false
}

func (bot Bot) trigger(eventName string, event *Event) {
	if callback, ok := bot.Callbacks[eventName]; ok {
		callback(event)
	}
}
