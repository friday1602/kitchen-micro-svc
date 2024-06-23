package common

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "failed to marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)

}

func ReadJSON(r *http.Request, data any) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func ResponseWithError(w http.ResponseWriter, status int, msg string) {
	ResponseWithJSON(w, status, map[string]string{"error": msg})
}
