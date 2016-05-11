package fbot

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Configure(Config{
		AccessToken: "",
		VerifyToken: "",
	})

	os.Exit(m.Run())
}

func TestWrongVerifyToken(t *testing.T) {
	Configure(Config{
		VerifyToken: "mytoken",
	})

	server := httptest.NewServer(Handler())

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
	Configure(Config{
		VerifyToken: "mytoken",
	})

	server := httptest.NewServer(Handler())

	defer server.Close()

	url := server.URL + "?hub.verify_token=mytoken&hub.challenge=challenge"

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
