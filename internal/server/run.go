package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func listenAndServe(logger *slog.Logger, server *http.Server) {
	logger.Info("listening on: " + server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("error listening and serving", err)
	}
}

func gracefulShutdown(
	ctx context.Context, logger *slog.Logger,
	server *http.Server, wg *sync.WaitGroup,
) {
	defer wg.Done()
	<-ctx.Done()

	shutdownCtx := context.Background()
	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("error shutting down http server", err)
	}
}

func run(ctx context.Context, logger *slog.Logger, server *http.Server) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	go listenAndServe(logger, server)

	var wg sync.WaitGroup
	wg.Add(1)

	go gracefulShutdown(ctx, logger, server, &wg)
	wg.Wait()

	return nil
}

func Run(ctx context.Context, logger *slog.Logger, server *http.Server,
) error {
	return run(ctx, logger, server)
}
