package server

import (
	"context"
	"net/http"
	"time"

	"github.com/quixsi/app/internal/config"
	"github.com/quixsi/app/internal/database"
	"github.com/quixsi/app/internal/server/handlers/users"
)

type Server struct {
	config      *config.Config
	server      *http.Server
	db          database.Database
	userHandler *users.Handler
}

func New(cfg *config.Config, db database.Database) *Server {
	mux := http.NewServeMux()

	s := &Server{
		config:      cfg,
		db:          db,
		userHandler: users.NewHandler(db),
		server: &http.Server{
			Addr:         cfg.Address,
			Handler:      mux,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}

	s.setupRoutes(mux)
	return s
}

func (s *Server) setupRoutes(mux *http.ServeMux) {
	// Add CORS middleware
	mux.HandleFunc("/", s.corsMiddleware(s.handleNotFound))
	
	// Health check route
	mux.HandleFunc("GET /health", s.corsMiddleware(s.handleHealth))
	
	// API v1 routes
	mux.HandleFunc("GET /api/v1/users", s.corsMiddleware(s.userHandler.GetUsers))
	mux.HandleFunc("POST /api/v1/users", s.corsMiddleware(s.userHandler.CreateUser))
	mux.HandleFunc("GET /api/v1/users/{id}", s.corsMiddleware(s.userHandler.GetUser))
	mux.HandleFunc("PUT /api/v1/users/{id}", s.corsMiddleware(s.userHandler.UpdateUser))
	mux.HandleFunc("DELETE /api/v1/users/{id}", s.corsMiddleware(s.userHandler.DeleteUser))
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

func (s *Server) handleNotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not found", http.StatusNotFound)
}

func (s *Server) corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", s.config.CORSOrigins[0])
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next(w, r)
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}