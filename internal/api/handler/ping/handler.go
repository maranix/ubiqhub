package ping

import "net/http"

func ApiPingHandler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				handlePing(w, r)
			} else {
				http.NotFoundHandler()
			}
		},
	)
}
