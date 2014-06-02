package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	ts := testServer(t)
	defer ts.Close()

	body := bytes.NewBufferString(`{"login":"brian","password":"p4$$"}`)
	resp, err := http.Post(ts.URL+"/auth", "application/json", body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != "logged in brian!" {
		t.Fatalf("expected to loggin as brian\ngot: %q", b)
	}
}

func testServer(t testing.TB) *httptest.Server {
	mux := http.NewServeMux()
	// TODO DRY logic
	mux.HandleFunc("/auth", Auth)
	return httptest.NewServer(mux)
}
