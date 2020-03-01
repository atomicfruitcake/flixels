package auth

import (
	"log"
	"net/http"
)

// Middleware middleware for checking adminToken in HTTP Header is valid
type Middleware struct {
	tokenUsers map[string]string
}

// Populate initialise the auth middleware and populate with the valid admin tokens
func (mw *Middleware) Populate() {
	if mw.tokenUsers == nil {
        mw.tokenUsers = make(map[string]string)
    }
	mw.tokenUsers["example"] = "adminToken"
}

// Middleware Function to apply to router the check all HTTP requests for admin tokens header
func (mw *Middleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("adminToken")
		if user, found := mw.tokenUsers[token]; found {
			log.Printf("Successfully Authenticated user %s\n", user)
			next.ServeHTTP(w, r)
		} else {
			log.Printf("Token %s is not a valid adminToken", token)
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
