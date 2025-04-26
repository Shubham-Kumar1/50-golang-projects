package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shubham-Kumar1/chat-app/config"
	"github.com/Shubham-Kumar1/chat-app/websocket"
)

type Server struct {
	config *config.Config
	pool   *websocket.Pool
	server *http.Server
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		config: cfg,
		pool:   websocket.NewPool(),
	}
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		return
	}

	client := websocket.NewClient(conn, s.pool)
	s.pool.Register <- client
	client.Read()
}

func (s *Server) setupRoutes() {
	http.HandleFunc(s.config.WebSocketPath, s.handleWebSocket)
}

func (s *Server) Start() error {
	s.setupRoutes()
	s.pool.Start()

	s.server = &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: nil,
	}

	// Channel to listen for errors coming from the server
	serverErrors := make(chan error, 1)

	// Start the server
	go func() {
		fmt.Println("Server started on port", s.config.Port)
		serverErrors <- s.server.ListenAndServe()
	}()

	// Channel to listen for an interrupt or terminate signal from the OS
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Blocking main and waiting for shutdown
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case <-shutdown:
		fmt.Println("Shutting down server...")

		// Stop the WebSocket pool
		s.pool.Stop()

		fmt.Println("Server shutdown complete")
		return nil
	}
}

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Create and start server
	server := NewServer(cfg)
	if err := server.Start(); err != nil {
		fmt.Println("Server failed:", err)
		os.Exit(1)
	}
}
