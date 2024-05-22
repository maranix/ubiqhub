package route

import (
	"net/http"

	"github.com/maranix/ubiqhub/cfg"
	"github.com/maranix/ubiqhub/internal/api/handler/healthz"
	"github.com/maranix/ubiqhub/internal/api/handler/ping"
)

func RegisterRoot(mux *http.ServeMux, _ *cfg.Config) {
	mux.Handle("/", http.NotFoundHandler())

	mux.Handle("/api/healthz", healthz.ApiHealthzHandler())
	mux.Handle("/api/ping", ping.ApiPingHandler())
}
