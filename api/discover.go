package api

import (
	"net/http"

	"github.com/baileyjm02/muzz-backend-technical-exercise/user"
)

// DiscoverUsers returns a random list of users.
func DiscoverUsers(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)

	user, err := user.GetUnswipedUsers(userID)
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}

	WriteJSON(w, user)
}
