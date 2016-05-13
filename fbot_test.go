package fbot

import "testing"

func TestOKEventTrigger(t *testing.T) {
	var ok bool

	bot := NewBot(Config{})

	bot.On(EventMessage, func(_ *Event) {
		ok = true
	})

	bot.trigger(&Event{Message: &Message{}})

	if !ok {
		t.Error("Event must be called")
	}
}
