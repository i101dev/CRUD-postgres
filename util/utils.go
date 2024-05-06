package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {

	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	data, err := json.Marshal(payload)

	if err != nil {
		fmt.Printf("failed to encode order to JSON: %+v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func ParseBody(r *http.Request, x interface{}) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalf("Error reading request body: %s", err)
		return
	}

	if err := json.Unmarshal([]byte(body), x); err != nil {
		log.Fatalf("Error unmarshalling JSON: %s", err)
		return
	}
}
