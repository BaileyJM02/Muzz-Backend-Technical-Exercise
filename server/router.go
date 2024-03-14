package server

import (
	"fmt"
	"net/http"

	"github.com/baileyjm02/muzz-backend-technical-exercise/api"
	"github.com/gorilla/mux"
)

// GetRouter returns a list of routes available and being listened too on the server
func GetRouter() *mux.Router {
	// Add the endpoint for engine payloads
	router := mux.NewRouter().StrictSlash(true)
	router.Use(AccessControlAllowMiddleware)

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	router.HandleFunc("/user/create", api.CreateUser).Methods("POST")
	router.HandleFunc("/login", api.LoginUser).Methods("POST")

	// Create a subrouter for authenticated routes, ensuring we're using the authenticated middleware
	authRouter := router.NewRoute().Subrouter()
	authRouter.Use(AuthenticatedMiddleware)

	authRouter.HandleFunc("/discover", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "discover")
	}).Methods("GET")

	return router
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.WriteHeader(404)
	fmt.Fprintf(w, "route not found.")
}
