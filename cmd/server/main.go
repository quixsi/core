package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/quixsi/app/internal/config"
	"github.com/quixsi/app/internal/database/inmemory"
	"github.com/quixsi/app/internal/server"
)

func main() {
	cfg := config.Load()
	db := inmemory.New()

	srv := server.New(cfg, db)

	go func() {
		log.Printf("Server starting on %s", cfg.Address)
		if err := srv.Start(); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}