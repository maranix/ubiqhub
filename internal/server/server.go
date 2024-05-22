package server

import (
	"log/slog"
	"net"
	"net/http"

	"github.com/maranix/ubiqhub/cfg"
	"github.com/maranix/ubiqhub/internal/api/middleware"
	"github.com/maranix/ubiqhub/internal/api/route"
)

func newServer(logger *slog.Logger, cfg *cfg.Config) *http.Server {
	mux := http.NewServeMux()

	route.RegisterRoot(mux, cfg)

	var handler http.Handler = mux

	handler = middleware.LoggingMiddleware(logger, handler)

	server := &http.Server{
		Addr:    net.JoinHostPort(cfg.Server.Host, cfg.Server.Port),
		Handler: handler,
	}

	return server
}

func CreateNewServer(logger *slog.Logger, cfg *cfg.Config) *http.Server {
	return newServer(logger, cfg)
}
