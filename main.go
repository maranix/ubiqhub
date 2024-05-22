package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/maranix/ubiqhub/cfg"
	"github.com/maranix/ubiqhub/internal/server"
)

func main() {
	stderr := os.Stderr
	ctx := context.Background()
	logger := slog.Default()
	config := cfg.Config{
		Server: cfg.Server{
			Host: "127.0.0.1",
			Port: "6969",
		},
	}

	srv := server.CreateNewServer(logger, &config)

	if err := server.Run(ctx, logger, srv); err != nil {
		fmt.Fprintf(stderr, "%s\n", err)
		os.Exit(1)
	}
}
