package main

import (
	"context"
	"flag"
	"fmt"
	"io"
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

func main() {
	w := os.Stderr
	ctx := context.Background()

	logger := createLogger(w)

	args, err := readEnvFile(env)
	if err != nil {
		fmt.Fprintf(w, "Could not read env file: %s\n", err)
		os.Exit(1)
	}
	config := cfg.FromArgs(args)

	srv := server.CreateNewServer(logger, config)
	if err := server.Run(ctx, logger, srv); err != nil {
		fmt.Fprintf(w, "%s\n", err)
		os.Exit(1)
	}
}

func readEnvFile(filenames ...string) (map[string]string, error) {
	args, err := godotenv.Read(filenames...)
	if err != nil {
		return nil, err
	}

	return args, nil
}

func createLogger(w io.Writer) *slog.Logger {
	logLevel := new(slog.LevelVar)
	handler := slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: logLevel,
	})

	slog.SetDefault(slog.New(handler))

	if releaseFlag {
		logLevel.Set(slog.LevelError)
	}

	return slog.Default()
}
