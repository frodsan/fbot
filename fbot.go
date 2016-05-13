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

// Event represents the event fired by the webhook.
type Event struct {
	Sender    Sender       `json:"sender"`
	Recipient Recipient    `json:"recipient"`
	Timestamp int64        `json:"timestamp,omitempty"`
	Message   *MessageInfo `json:"message"`
	Delivery  *Delivery    `json:"delivery"`
	Postback  *Postback    `json:"postback"`
}

// Sender represents the user who sent the message.
type Sender struct {
	ID int64 `json:"id"`
}

// Recipient represents the user who the message was sent to.
type Recipient struct {
	ID int64 `json:"id"`
}

// MessageInfo represents the message callback object.
type MessageInfo struct {
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
