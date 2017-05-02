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
	queryParams := url.Values{"hub.mode": {"invalid"}}

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

func TestWrongVerifyToken(t *testing.T) {
	queryParams := url.Values{
		"hub.mode":         {"subscribe"},
		"hub.verify_token": {"wrong"},
	}

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

func TestOKVerifyToken(t *testing.T) {
	queryParams := url.Values{
		"hub.mode":         {"subscribe"},
		"hub.verify_token": {"ok"},
		"hub.challenge":    {"challenge"},
	}

	req, error := http.NewRequest("GET", "/webhook?"+queryParams.Encode(), nil)

	if error != nil {
		t.Fatal(error)
	}

	bot := NewBot()
	bot.VerifyToken = "ok"

	res := httptest.NewRecorder()

	handler := Handler(bot)

	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v; got %v", http.StatusOK, status)
	}

	expected := queryParams["hub.challenge"][0]

	if string(res.Body.Bytes()) != expected {
		t.Errorf("Expected %s, got %s", expected, res.Body.Bytes())
	}
}
