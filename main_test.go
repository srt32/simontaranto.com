package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUnknownURLRedirect(t *testing.T) {
	requestedURL := "http://simontaranto.com/nonsense"
	expectedURL := "http://simontaranto.com"

	handler := RedirectHandler
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", requestedURL, nil)
	assert.Nil(t, err)

	handler.ServeHTTP(w, req)
}
