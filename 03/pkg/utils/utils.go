package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		return
	}

	err = json.Unmarshal(body, x)
	if err != nil {
		log.Printf("Error unmarshaling request body: %v", err)
	}
}
