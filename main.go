package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/bootstrap"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/pkg/config"
	"github.com/labstack/gommon/log"
)

func main() {
	app := bootstrap.Bootstrap()
	if app == nil {
		return
	}
	cfg := config.GetConfig()

	// Set up signal handling
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start the server in a goroutine so we can listen for shutdown signals
	go func() {
		if err := app.Start(fmt.Sprintf(":%d", cfg.AppPort)); err != nil &&
			err != http.ErrServerClosed {
			log.Fatal("Error starting server:", err)
		}
	}()

	// Wait for an interrupt signal (Ctrl+C) to shut down the server
	<-ctx.Done()
	log.Info("Received shutdown signal, initiating graceful shutdown...")

	// Set a timeout for the graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := app.Shutdown(ctx); err != nil {
		log.Fatal("Failed to gracefully shut down the server:", err)
	} else {
		log.Info("Server gracefully shut down")
	}
}
