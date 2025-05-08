package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type helloResponse struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(helloResponse{Message: "Hello, World!"}); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", http.HandlerFunc(helloHandler)))
}
