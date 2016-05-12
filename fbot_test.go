package fbot

import "testing"

func TestPanicIfEventNameIsInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Given event name must be invalid")
		}
	}()

	bot := NewBot(Config{})

	bot.On("invalid", func(_ *Event) {})
}

func TestOKEventTrigger(t *testing.T) {
	var ok bool

	bot := NewBot(Config{})

	bot.On("message", func(_ *Event) {
		ok = true
	})

	bot.trigger("message", &Event{})

	if !ok {
		t.Error("Event must be called")
	}
}
