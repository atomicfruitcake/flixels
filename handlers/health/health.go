package health

import (
	"net/http"
)

// Handler HTTP Request handler for health checks
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
