package logging

import (
	"log"
	"net/http"
)

// Middleware middleware logging requests handled by a router
func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Received %s request to %s",  r.Method, r.RequestURI)
        next.ServeHTTP(w, r)
    })
}