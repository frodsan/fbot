package fbot

import "testing"

func TestOKEventTrigger(t *testing.T) {
	var ok bool

	bot := NewBot()

	bot.On("event", func() {
		ok = true
	})

	bot.Trigger("event")

	if !ok {
		t.Error("Event 'event' must be triggered")
	}
}
