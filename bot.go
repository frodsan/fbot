package fbot

type Bot struct {
	AccessToken string
	AppSecret   string
	VerifyToken string
	Callbacks   map[string]func()
}

func NewBot() Bot {
	return Bot{
		Callbacks: make(map[string]func()),
	}
}

func (bot Bot) On(event string, callback func()) {
	bot.Callbacks[event] = callback
}

func (bot Bot) Trigger(event string) {
	if callback, ok := bot.Callbacks[event]; ok {
		callback()
	}
}
