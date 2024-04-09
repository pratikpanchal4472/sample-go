package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pratikpanchal4472/common/component/utils"
)

type Server struct {
	ListenPort int
	Router     *gin.Engine
}

func (s *Server) GetListenPort() int {
	return s.ListenPort
}

func (s *Server) GetRouter() *gin.Engine {
	return s.Router
}

func NewServer(listenPort int) *Server {
	server := &Server{
		ListenPort: listenPort,
		Router:     utils.GetGinEngine(),
	}
	server.bindHandlers()
	return server
}

func (s *Server) bindHandlers() {
	s.Router.GET("/ping", s.PingHandler)
}

func main() {
	NewServer(8080).Start()
}

func (s *Server) Start() {
	utils.Start(s)
}

func (s *Server) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
