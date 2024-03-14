package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("aSuperSecretKeyStoredInAProperPlace")

// TokenResponse is the response for a token.
type TokenResponse struct {
	Token        string    `json:"token"`
	ExpiresAfter time.Time `json:"expires_after"`
}

// CreateTokenString creates a token string for a user using JWT.
func CreateTokenString(email string) (TokenResponse, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &jwt.StandardClaims{
		Subject:   email,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return TokenResponse{}, err
	}

	return TokenResponse{
		Token:        tokenString,
		ExpiresAfter: expirationTime,
	}, nil
}

// ParseToken parses a token string and returns the token - checking for expiry.
func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, fmt.Errorf("token is expired")
	}

	return token, nil
}
