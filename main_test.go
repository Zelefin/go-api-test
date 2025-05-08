package main

import (
	"io"
	"strings"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	res := w.Result()
	defer func() {
		if err := res.Body.Close(); err != nil {
			t.Errorf("error closing response body: %v", err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		t.Errorf("wanted status OK, but got %v", res.StatusCode)
	}

	body, err := io.ReadAll(w.Body)

	if err != nil {
		t.Errorf("error reading body: %v", err)
	}

	want := `{"message":"Hello, World!"}`
	got := strings.TrimSpace(string(body))

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}
