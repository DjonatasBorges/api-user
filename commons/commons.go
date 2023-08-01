package commons

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleServerError(w http.ResponseWriter, err error) {
	log.Println("Error:", err)
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		HandleServerError(w, err)
	}
}
