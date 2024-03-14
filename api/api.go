package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("API Starting!")
}

func Start() {
	fmt.Println("API Started!")

}

func WriteErrorJSON(w http.ResponseWriter, err error) {
	errorMessage := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}

	WriteJSON(w, errorMessage)
}

func WriteJSON(w http.ResponseWriter, any interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	bytes, err := json.Marshal(any)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error marshalling response: %s", err)
		return
	}

	w.Write(bytes)
}
