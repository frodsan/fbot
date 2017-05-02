package fbot

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestMethodNotAllowed(t *testing.T) {
	req, eres := http.NewRequest("HEAD", "/webhook", nil)

	if eres != nil {
		t.Fatal(eres)
	}

	bot := NewBot()

	res := httptest.NewRecorder()

	handler := Handler(bot)

	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %v; got %v", http.StatusMethodNotAllowed, status)
	}
}

func TestWrongVerifyMode(t *testing.T) {
	queryParams := url.Values{"mode": {"invalid"}}

	req, eres := http.NewRequest("GET", "/webhook?"+queryParams.Encode(), nil)

	if eres != nil {
		t.Fatal(eres)
	}

	bot := NewBot()

	res := httptest.NewRecorder()

	handler := Handler(bot)

	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusForbidden {
		t.Errorf("Expected status code %v; got %v", http.StatusForbidden, status)
	}
}
