package fbot

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethodNotAllowed(t *testing.T) {
	bot := NewBot(Config{})
	server := httptest.NewServer(Handler(bot))

	defer server.Close()

	res, err := http.Head(server.URL)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Error("HTTP method must be not allowed")
	}
}

func TestWrongVerifyToken(t *testing.T) {
	bot := NewBot(Config{
		VerifyToken: "token",
	})

	server := httptest.NewServer(Handler(bot))

	defer server.Close()

	url := server.URL + "?hub.verify_token=wrongtoken&hub.challenge=challenge"

	res, err := http.Get(url)

	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	if string(body) == "challenge" {
		t.Errorf("Expected error; got challenge: %s", body)
	}
}

func TestOKVerifyToken(t *testing.T) {
	bot := NewBot(Config{
		VerifyToken: "token",
	})

	server := httptest.NewServer(Handler(bot))

	defer server.Close()

	url := server.URL + "?hub.verify_token=token&hub.challenge=challenge"

	res, err := http.Get(url)

	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "challenge" {
		t.Errorf("Expected 'challenge'; got '%s'", body)
	}
}

func TestReceiveWithEmptySignature(t *testing.T) {
	bot := NewBot(Config{})

	server := httptest.NewServer(Handler(bot))

	defer server.Close()

	var json []byte

	res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(json))

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status %d; got %d", http.StatusBadRequest, res.StatusCode)
	}
}

func TestReceiveMessage(t *testing.T) {
	bot := NewBot(Config{})

	server := httptest.NewServer(Handler(bot))

	defer server.Close()

	json := `{"object":"page","entry":[{"id":"1726989104239153","time":1463564888322,"messaging":[{"sender":{"id":"117842461964948"},"recipient":{"id":"1726989104239153"},"timestamp":1463563811474,"message":{"mid":"mid.1463563811462:9bff5c416e86887852","seq":8,"text":"hio"}}]}]}`

	res, err := http.Post(server.URL, "application/json", bytes.NewBuffer([]byte(json)))

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status %d; got %d", http.StatusBadRequest, res.StatusCode)
	}

}
