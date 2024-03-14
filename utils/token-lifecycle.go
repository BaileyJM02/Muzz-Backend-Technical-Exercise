package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("aSuperSecretKeyStoredInAProperPlace")

// TokenResponse is the response for a token.
type TokenResponse struct {
	UserID       int       `json:"user_id"`
	Token        string    `json:"token"`
	ExpiresAfter time.Time `json:"expires_after"`
}

// CreateTokenString creates a token string for a user using JWT.
func CreateTokenString(id int) (TokenResponse, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &jwt.StandardClaims{
		Id:        fmt.Sprintf("%d", id),
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
func ParseToken(tokenString string) (TokenResponse, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return TokenResponse{}, err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	if claims.ExpiresAt < time.Now().Unix() {
		return TokenResponse{}, fmt.Errorf("token is expired")
	}

	userID, err := strconv.Atoi(claims.Id)
	if err != nil {
		return TokenResponse{}, err
	}

	return TokenResponse{
		UserID:       userID,
		Token:        tokenString,
		ExpiresAfter: time.Unix(claims.ExpiresAt, 0),
	}, nil
}
