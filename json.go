package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func errResponse(w http.ResponseWriter, code int, errMessage string) {
	if code > 499 {
		log.Printf("Returning a 5XX error: %v", errMessage)
	}
	type errResponderStruct struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponderStruct{
		Error: errMessage,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Could not marshal json response: %v", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
