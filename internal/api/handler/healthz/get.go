package healthz

import "net/http"

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.RequestURI))
}
