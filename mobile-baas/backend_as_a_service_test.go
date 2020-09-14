package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestDefaultHandler(t *testing.T) {
	payload := strings.NewReader("")
	r := httptest.NewRequest("GET", "/", payload)
	w := httptest.NewRecorder()
	defaultHandler(w, r)
	resp := w.Result()
	expected := http.StatusOK
	result := resp.StatusCode
	if result != expected {
		t.Errorf("TestHandleDefault: Expected: %d, got %d\n", expected, result)
	}
}

func TestTimeHandler(t *testing.T) {
	now := time.Now()
	payload := strings.NewReader("")
	r := httptest.NewRequest("GET", "/time", payload)
	w := httptest.NewRecorder()
	timeHandler(w, r)
	resp := w.Result()
	expected := http.StatusOK
	result := resp.StatusCode
	if result != expected {
		t.Errorf("TestHandleDefault: Expected: %d, got %d\n", expected, result)
	}

	fmt.Println("Got time ", w.Body.String(), ", our time is ", now.String())
}
