package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Status string `json:"status"`
	Body string `json:"body"`
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := Message{
		Status: "success",
		Body: "Request processed successfully",
	}
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/ping", rateLimiter(endpointHandler))
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
