package fbot

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
