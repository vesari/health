package health

import (
	"encoding/json"
	"net/http"
)

// Handler is a HTTP Server Handler implementation
type Handler struct {
	CompositeChecker
}

// ServeHTTP returns a json encoded Health
// set the status to http.StatusServiceUnavailable if the check is down
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	health := h.CompositeChecker.Check()

	if health.IsDown() {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}
