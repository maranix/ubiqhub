package healthz

import "net/http"

func Handler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				handleGet(w, r)
			} else {
				http.NotFoundHandler()
			}
		},
	)
}
