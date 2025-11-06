package main

import (
	"encoding/json"
	"log"
	"net/http"
	tollbooth "github.com/didip/tollbooth/v7"
)


type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := Message{
		Status: "success",
		Body:   "Request processed successfully",
	}
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Println(err)
		return
	}
}


func main() {
	message := Message{
				Status: "Request Failed",
				Body:  "Rate limit exceeded. Please try again later.",
	}
	jsonMessage, _ := json.Marshal(message)
	tollboothLimiter := tollbooth.NewLimiter(1,nil)
	tollboothLimiter.SetMessage(string(jsonMessage))
	tollboothLimiter.SetMessageContentType("application/json")

	http.Handle("/ping", tollbooth.LimitFuncHandler(tollboothLimiter, endpointHandler))
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}