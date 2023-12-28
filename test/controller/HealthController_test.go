package controller_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"star/src/router"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := router.InitRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}
