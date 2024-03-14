package server

import (
	"context"
	"net/http"

	"github.com/baileyjm02/muzz-backend-technical-exercise/utils"
)

const (
	AccessControlAllowOrigin  = "*"
	AccessControlAllowMethods = "GET" // "POST, GET, OPTIONS, PUT, DELETE"
	AccessControlAllowHeaders = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header"
)

// AccessControlAllowMiddleware headers
func AccessControlAllowMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AccessControlAllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", AccessControlAllowMethods)
		w.Header().Set("Access-Control-Allow-Headers", AccessControlAllowHeaders)
		next.ServeHTTP(w, r)
	})
}

// AuthenticatedMiddleware ensures that the user is authenticated before allowing access to the route.
// Accepts a token in the Authorization header or as a cookie.
func AuthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		cookieToken, _ := r.Cookie("token")
		if authHeader == "" && cookieToken == nil {
			unauthorisedResponse(w)
			return
		}

		var token string
		if authHeader != "" {
			token = authHeader
		} else {
			token = cookieToken.Value
		}

		parsedToken, err := utils.ParseToken(token)
		if err != nil {
			unauthorisedResponse(w)
			return
		}

		// Add the UserID to the context of the next request
		ctx := context.WithValue(r.Context(), "userID", parsedToken.UserID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// unauthorisedResponse returns a 401 response to the client.
func unauthorisedResponse(w http.ResponseWriter) {
	w.WriteHeader(401)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"error": "invalid token"}`))
}
