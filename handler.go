package health

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	CompositeChecker
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	health := h.CompositeChecker.Check()

	if health.IsDown() {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}
