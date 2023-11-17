package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingEndpoint(t *testing.T) {
	// Create a new Gin router
	r := gin.New()

	// Set the router to release mode
	gin.SetMode(gin.ReleaseMode)

	// Define the /ping route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Create a new HTTP request to the /ping endpoint
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	recorder := httptest.NewRecorder()

	// Use the router to handle the request
	r.ServeHTTP(recorder, req)

	// Check if the status code is 200 OK
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Check if the response body contains the expected message
	expected := `{"message":"pong"}`
	assert.Equal(t, expected, recorder.Body.String())
}
