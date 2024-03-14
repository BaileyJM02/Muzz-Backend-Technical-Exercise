package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/baileyjm02/muzz-backend-technical-exercise/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := user.CreateRandom()

	// Send user as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	bytes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error marshalling user: %s", err)
		return
	}

	w.Write(bytes)
}
