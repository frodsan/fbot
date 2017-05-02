package fbot

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestMethodNotAllowed(t *testing.T) {
	req, error := http.NewRequest("HEAD", "/webhook", nil)

	if error != nil {
		t.Fatal(error)
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

	req, error := http.NewRequest("GET", "/webhook?"+queryParams.Encode(), nil)

	if error != nil {
		t.Fatal(error)
	}

	bot := NewBot()

	res := httptest.NewRecorder()

	handler := Handler(bot)

	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusForbidden {
		t.Errorf("Expected status code %v; got %v", http.StatusForbidden, status)
	}
}
