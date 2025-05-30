package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Server represents the HTTP server
type Server struct {
	router                *http.ServeMux
	server                *http.Server
	handlerPackageManager *HandlerPacksManager
}

// Config holds the server configuration
type Config struct {
	Port            int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

// NewServer creates a new HTTP server
func NewServer(config Config, h *HandlerPacksManager) *Server {

	// Create router
	r := http.NewServeMux()

	registerRoutes(r, h)

	// Create server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: r,
	}

	return &Server{
		router:                r,
		server:                srv,
		handlerPackageManager: h,
	}
}

// ListenAndServe starts the HTTP server and listen on specified TCP port.
func (s *Server) ListenAndServe() error {
	return s.server.ListenAndServe()
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

// DefaultConfig returns the default server configuration
func DefaultConfig() Config {
	return Config{
		Port:            8080,
		ReadTimeout:     15 * time.Second,
		WriteTimeout:    15 * time.Second,
		ShutdownTimeout: 5 * time.Second,
	}
}
