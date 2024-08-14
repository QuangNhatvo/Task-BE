package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// method GET
func TestResponse(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
func TestUservalue(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/foo?value=bar", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	Expected := `{"user":"foo","value":"bar"}`
	assert.Equal(t, Expected, w.Body.String())
}

// method POST
func TestAuthorized(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin", bytes.NewBufferString(`{"value": "new_value"}`))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Basic Zm9vOmJhcg==")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	Expected := `{"status":"ok"}`
	assert.Equal(t, Expected, w.Body.String())
}
