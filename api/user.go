package api

import (
	"errors"
	"net/http"

	"github.com/baileyjm02/muzz-backend-technical-exercise/user"
	"github.com/baileyjm02/muzz-backend-technical-exercise/utils"
)

// CreateUser creates a new user and returns it.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := user.CreateRandom()
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}

	WriteJSON(w, ResultWrapper{
			Result: user,
		})
}

// LoginUser attempts to authenticate a user and return a token. It specifically returns the same
// error message for invalid credentials to avoid leaking information about the existence of a user.
func LoginUser(w http.ResponseWriter, r *http.Request) {
	errorResponse := errors.New("invalid credentials")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := user.GetByEmail(email)
	if err != nil {
		WriteErrorJSON(w, errorResponse)
		return
	}

	if !utils.CompareHashPassword(password, user.Password) {
		WriteErrorJSON(w, errorResponse)
		return
	}

	token, err := utils.CreateTokenString(user.ID)
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}

	// Set the token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token.Token,
		Expires: token.ExpiresAfter,
	})

	WriteJSON(w, token)
}
