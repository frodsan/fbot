package fbot

import "testing"

func TestPanicIfEventNameIsInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Given event name must be invalid")
		}
	}()

	On("invalid", func(_ Event) {})
}
