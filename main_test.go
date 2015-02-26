package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUnknownURLRedirect(t *testing.T) {
	requestedURL := "http://simontaranto.com/nonsense"
	//expectedURL := "http://simontaranto.com"

	handler := http.HandlerFunc(RedirectHandler)
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", requestedURL, nil)
	assert.Nil(t, err)

	handler.ServeHTTP(w, req)
	if w.Code != 301 {
		t.Errorf("Response was not 301, was: %v", w.Code)
	}

	//if w.Path != expectedURL {
	//	t.Errorf("Redirected URL was not %v, was: %v", expectedURL, w.Path)
	//}
}
