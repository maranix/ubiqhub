package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/maranix/ubiqhub/cfg"
	"github.com/maranix/ubiqhub/internal/api/handler/healthz"
	"github.com/maranix/ubiqhub/internal/api/middleware"
)

func newServer(logger *slog.Logger, cfg *cfg.Config) *http.Server {
	mux := http.NewServeMux()
	addRoutes(mux, cfg)

	var handler http.Handler = mux

	handler = middleware.LoggingMiddleware(logger, handler)

	server := &http.Server{
		Addr:    net.JoinHostPort("127.0.0.1", "6969"),
		Handler: handler,
	}

	return server
}

func addRoutes(mux *http.ServeMux, _ *cfg.Config) {
	mux.Handle("/api/healthz", healthz.Handler())
	mux.Handle("/", http.NotFoundHandler())
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	config := cfg.Config{
		Server: cfg.Server{
			Host: "127.0.0.1",
			Port: "6969",
		},
	}

	logger := slog.Default()
	server := newServer(logger, &config)

	go func() {
		logger.Info("listening on: " + server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("error listening and serving", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()

		// make a new context for the Shutdown (thanks Alessandro Rosetti)
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Error("error shutting down http server", err)
		}
	}()
	wg.Wait()

	return nil
}

func main() {
	ctx := context.Background()
	stderr := os.Stderr

	if err := run(ctx); err != nil {
		fmt.Fprintf(stderr, "%s\n", err)
		os.Exit(1)
	}
}
