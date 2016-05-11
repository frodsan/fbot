package fbot

// Config represents the Bot configuration options.
type Config struct {
	AccessToken string
	AppSecret   string
	VerifyToken string
}

var config Config

// Configure sets Bot configuration options.
func Configure(c Config) {
	config = c
}
