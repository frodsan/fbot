package fbot

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("HEAD", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	bot := NewBot()

	rr := httptest.NewRecorder()

	handler := Handler(bot)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v expected %v", status, http.StatusMethodNotAllowed)
	}
}
