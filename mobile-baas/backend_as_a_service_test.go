package mbaas

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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

// func TestGetTimeHandler(t *testing.T) {
// 	r := httptest.NewRequest("GET", "http://time", nil)
// 	w := httptest.NewRecorder()
// 	defaultHandler(w, r)
// 	resp := w.Result()
// 	expected := http.StatusOK
// 	result := resp.StatusCode
// 	if result != expected {
// 		t.Errorf("TestHandleDefault: Expected: %d, got %d\n", expected, result)
// 	}
// }
