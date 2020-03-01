package root

import (
	"fmt"
	"net/http"

)

// Handler HTTP Request handler for requests to server root
func Handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("http://%s/health", r.Host), 302)
}
