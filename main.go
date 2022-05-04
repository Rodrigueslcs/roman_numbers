package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Rodrigueslcs/roman_numbers/roman_numbers"
)

func main() {
	// s := "AXDxBLX"
	// roman_numbers.Roman(s)
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ": 8080",
		ErrorLog:     logger,
	}
	http.HandleFunc("/search", search)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func search(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		errorMessage := "erro request"
		var request roman_numbers.Request
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		response := roman_numbers.Roman(request.Text)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}

	default:
		fmt.Fprintf(w, "METHOD INVALID")

	}
}
