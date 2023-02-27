package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetHealthz(t *testing.T) {
	// Create a new Gin router instance
	r := gin.Default()
	// Create a new HealthzController instance and register the routes
	NewHealthzController(r)
	// Create a new HTTP request to the health check endpoint
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a new HTTP response recorder
	recorder := httptest.NewRecorder()
	// Serve the HTTP request to the router
	r.ServeHTTP(recorder, req)
	// Check that the HTTP response code is 200 OK
	assert.Equal(t, http.StatusOK, recorder.Code)
	// Check that the HTTP response body is {"status": "OK"}
	assert.Equal(t, "{\"status\":\"OK\"}", recorder.Body.String())
}
