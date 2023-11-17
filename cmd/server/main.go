package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	ListenPort int
	Router     *gin.Engine
}

func NewServer(listenPort int) *Server {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	server := &Server{
		ListenPort: listenPort,
		Router:     engine,
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
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0%d", s.ListenPort),
		Handler: s.Router,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic("Server Start Error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Println("Shutdown failed")
	}
	<-shutdownCtx.Done()
	log.Println("Server Existing")
}

func (s *Server) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
