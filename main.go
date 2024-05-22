package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/maranix/ubiqhub/cfg"
	"github.com/maranix/ubiqhub/internal/server"
)

func main() {
	stderr := os.Stderr
	ctx := context.Background()
	logger := slog.Default()

	args, err := godotenv.Read()
	if err != nil {
		fmt.Fprintf(stderr, "Could not read env variable: %s\n", err)
		os.Exit(1)
	}

	config := cfg.FromArgs(args)
	srv := server.CreateNewServer(logger, config)

	if err := server.Run(ctx, logger, srv); err != nil {
		fmt.Fprintf(stderr, "%s\n", err)
		os.Exit(1)
	}
}
