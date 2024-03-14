package api

import (
	"net/http"
	"strconv"

	"github.com/baileyjm02/muzz-backend-technical-exercise/user"
)

// DiscoverUsers returns a random list of users.
func DiscoverUsers(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)

	// Get request parameters from URL query
	// Not bothering to check for errors here, as if it's invalid, we're just going to ignore it.
	urlParams := r.URL.Query()
	minAgeStr := urlParams.Get("minAge")
	minAge, _ := strconv.Atoi(minAgeStr)

	maxAgeStr := urlParams.Get("maxAge")
	maxAge, _ := strconv.Atoi(maxAgeStr)

	filters := user.DiscoverFilters{
		MinAge: minAge,
		MaxAge: maxAge,
		Gender: urlParams.Get("gender"),
	}

	user, err := user.GetUnswipedUsers(userID, filters)
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}

	WriteJSON(w, user)
}
