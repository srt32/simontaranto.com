package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUnknownURLRedirect(t *testing.T) {
	requestedURL := "http://simontaranto.com/nonsense"
	expectedURL := "http://blog.simontaranto.com"

	handler := http.HandlerFunc(RedirectHandler)
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", requestedURL, nil)
	if err != nil {
		t.Errorf("error making request: %v", err)
	}

	handler.ServeHTTP(w, req)
	if w.Code != 301 {
		t.Errorf("Response was not 301, was: %v", w.Code)
	}

	actualURL := w.HeaderMap["Location"][0]
	if actualURL != expectedURL {
		t.Errorf("Redirected URL was not %v, was: %v", expectedURL, actualURL)
	}
}

func TestKnownURLRedirect(t *testing.T) {
	requestedURL := "http://simontaranto.com/2015/01/11/datetime-parsing-with-go.html?foo=bar"
	expectedURL := "http://blog.simontaranto.com/post/2015-01-11-datetime-parsing-with-go.html"

	handler := http.HandlerFunc(RedirectHandler)
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", requestedURL, nil)
	if err != nil {
		t.Errorf("error making request: %v", err)
	}

	handler.ServeHTTP(w, req)
	if w.Code != 301 {
		t.Errorf("Response was not 301, was: %v", w.Code)
	}

	actualURL := w.HeaderMap["Location"][0]
	if actualURL != expectedURL {
		t.Errorf("Redirected URL was not %v, was: %v", expectedURL, actualURL)
	}

	actualReferer := w.HeaderMap["Referer"][0]
	if actualReferer != "simontaranto.com" {
		t.Errorf("Referer was not %v, was: %v", expectedURL, actualURL)
	}

	actualReferrer := w.HeaderMap["Referrer"][0]
	if actualReferrer != "simontaranto.com" {
		t.Errorf("Referer was not %v, was: %v", expectedURL, actualURL)
	}
}
