package main

import (
	"log"
	"net/http"
	"encoding/json"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Resposta sendo um erro 5XX:", msg)
	}
	type errResponse struct {
		Error string "json:'error'"
	}
		respondWithJSON(w, code, errResponse {
			Error: msg,
		})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Falha na resposta do JSON: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}