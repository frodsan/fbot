package fbot

type Bot struct {
	AccessToken string
	AppSecret   string
	VerifyToken string
	Callbacks   map[string]func()
}

func NewBot(options ...func(*Bot)) Bot {
	bot := Bot{
		Callbacks: make(map[string]func()),
	}

	for _, option := range options {
		option(&bot)
	}

	return bot
}

func (bot Bot) On(event string, callback func()) {
	bot.Callbacks[event] = callback
}

func (bot Bot) Trigger(event string) {
	if callback, ok := bot.Callbacks[event]; ok {
		callback()
	}
}
