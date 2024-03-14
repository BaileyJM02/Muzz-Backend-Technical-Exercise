package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Add singular and plural wrapper for JSON responses.
// Forgot to add these in the original implementation, was an afterthought.
type ResultWrapper struct {
	Result interface{} `json:"result"`
}

type ResultsWrapper struct {
	Results interface{} `json:"results"`
}

// WriteErrorJSON writes an error message to the response, in JSON format.
func WriteErrorJSON(w http.ResponseWriter, err error) {
	errorMessage := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	bytes, err := json.Marshal(errorMessage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error marshalling response: %s", err)
		return
	}

	w.Write(bytes)
}

// WriteJSON writes an interface to the response, in JSON format.
func WriteJSON(w http.ResponseWriter, any interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, err := json.Marshal(any)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error marshalling response: %s", err)
		return
	}

	w.Write(bytes)
}
