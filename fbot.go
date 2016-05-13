package fbot

// Version is the version of the library.
const Version = "0.0.1"

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

// On registers a callback for the given eventName.
func (bot *Bot) On(eventName string, callback func(*Event)) {
	bot.Callbacks[eventName] = callback
}

func (bot Bot) trigger(event *Event) {
	eventName := selectEvent(event)

	if callback, ok := bot.Callbacks[eventName]; ok {
		callback(event)
	}
}

func selectEvent(event *Event) string {
	var eventName string

	if event.Message != nil {
		eventName = EventMessage
	} else if event.Delivery != nil {
		eventName = EventDelivery
	} else if event.Postback != nil {
		eventName = EventPostback
	}

	return eventName
}
