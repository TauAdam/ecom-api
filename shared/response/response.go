package response

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("no body provided")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}
func SendJSON(w http.ResponseWriter, statusCode int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(payload)
}

func SendError(w http.ResponseWriter, statusCode int, errMessage error) {
	err := SendJSON(w, statusCode, map[string]string{"error": errMessage.Error()})
	if err != nil {
		log.Fatalf("Error writing response: %v", err)
	}
}
