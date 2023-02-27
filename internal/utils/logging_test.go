package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetClientIP(t *testing.T) {
	// Create a new Gin router instance
	r := gin.Default()
	// Add a handler function that calls the GetClientIP function
	r.GET("/test", func(c *gin.Context) {
		ip := GetClientIP(c)
		c.JSON(http.StatusOK, gin.H{"ip": ip})
	})
	// Create a new HTTP request to the test endpoint
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Set the X-Forwarded-For header to a fake IP address
	req.Header.Set("X-Forwarded-For", "10.0.0.1")
	// Create a new HTTP response recorder
	recorder := httptest.NewRecorder()
	// Serve the HTTP request to the router
	r.ServeHTTP(recorder, req)
	// Check that the HTTP response code is 200 OK
	assert.Equal(t, http.StatusOK, recorder.Code)
	// Check that the HTTP response body contains the correct IP address
	assert.Contains(t, recorder.Body.String(), "\"ip\":\"10.0.0.1\"")

	// Create a new HTTP request with an empty X-Forwarded-For header and no other headers
	reqEmpty, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a new HTTP response recorder
	recorderEmpty := httptest.NewRecorder()
	// Serve the HTTP request to the router
	r.ServeHTTP(recorderEmpty, reqEmpty)
	// Check that the HTTP response code is 200 OK
	assert.Equal(t, http.StatusOK, recorderEmpty.Code)
	// Check that the HTTP response body contains the remote IP address
	assert.Contains(t, recorderEmpty.Body.String(), "{\"ip\":\"\"}")
}

func TestGetDurationInMilliseconds(t *testing.T) {
	// Call the function to get the duration
	start := time.Now()
	time.Sleep(100 * time.Millisecond)
	duration := GetDurationInMillseconds(start)
	// Check that the duration is roughly 100 milliseconds
	assert.InDelta(t, 100.0, duration, 0.5)
}
