package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWithValidName(t *testing.T) {
	for _, name := range []string{"Ada", "Grace"} {
		req := httptest.NewRequest(http.MethodGet, "/hello?name="+name, nil)
		rec := httptest.NewRecorder()

		helloHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("name=%q: got status %d, want 200", name, rec.Code)
		}
		var got Greeting
		if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
			t.Fatalf("name=%q: invalid JSON body: %v", name, err)
		}
		if got.Name != name || got.Message == "" {
			t.Fatalf("name=%q: got greeting %+v", name, got)
		}
	}
}

func TestHelloWithMissingName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	helloHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("got status %d, want 400", rec.Code)
	}
	var got ErrorResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("invalid JSON body: %v", err)
	}
	if got.Code != http.StatusBadRequest || got.Message == "" {
		t.Fatalf("got error %+v", got)
	}
}

func TestHelloWithEmptyName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=", nil)
	rec := httptest.NewRecorder()

	helloHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("got status %d, want 400", rec.Code)
	}
}
