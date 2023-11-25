package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func SetupTestServer() *Server {
	return NewServer(8080)
}

func TestServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	server := SetupTestServer()

	t.Run("Test Request for /Ping", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/ping", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		server.PingHandler(c)

		expectedBody := `{"message":"pong"}`
		var expected, actual map[string]interface{}
		json.Unmarshal([]byte(expectedBody), &expected)
		json.Unmarshal(w.Body.Bytes(), &actual)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected Response %s but got %s", expectedBody, w.Body.String())
		}
		t.Log("Test Request for /Ping - SUCCESS")
	})
}
