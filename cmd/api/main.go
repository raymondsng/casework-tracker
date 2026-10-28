package main

import (
	"log"
	"net/http"

	"github.com/raymondsng/casework-tracker/internal/httpapi"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: httpapi.NewRouter(),
	}

	log.Println("API listening on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
