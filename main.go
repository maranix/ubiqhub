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

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = "0.0.0.0"
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "6969"
	}

	config := cfg.Config{
		Server: cfg.Server{
			Host: host,
			Port: port,
		},
	}

	srv := server.CreateNewServer(logger, &config)

	if err := server.Run(ctx, logger, srv); err != nil {
		fmt.Fprintf(stderr, "%s\n", err)
		os.Exit(1)
	}
}
