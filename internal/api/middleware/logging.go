package middleware

import (
	"log/slog"
	"net/http"
)

func loggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger.Info(r.Method + " " + r.URL.Path)
			next.ServeHTTP(w, r)
		},
	)
}

func LoggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return loggingMiddleware(logger, next)
}
