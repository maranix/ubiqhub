package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/maranix/ubiqhub/cfg"
	"github.com/maranix/ubiqhub/internal/server"
)

var (
	env         string
	releaseFlag bool
)

func init() {
	flag.BoolVar(&releaseFlag, "release", false, "runs the server in release/production mode")
	flag.Parse()

	if releaseFlag {
		env = ".env"
	} else {
		env = ".env.development"
	}
}

func readEnvFile(filenames ...string) (map[string]string, error) {
	args, err := godotenv.Read(filenames...)
	if err != nil {
		return nil, err
	}

	return args, nil
}

func main() {
	ctx := context.Background()
	logger := slog.Default()

	args, err := readEnvFile(env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read env file: %s\n", err)
		os.Exit(1)
	}
	config := cfg.FromArgs(args)

	srv := server.CreateNewServer(logger, config)
	if err := server.Run(ctx, logger, srv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
