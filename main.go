package main

import (
	"log"
	"net/http"
	"os"

	"go-learn-rest-api/handlers"
)

// How to try it: go run main.go
// How to test configuration env PORT:
// first: export PORT=8000 && go install go-learn-rest-api && ../../bin/go-learn-rest-api
// then: export PORT=8000 && ../../bin/go-learn-rest-api
func main() {
	log.Print("Starting the service...")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	router := handlers.Router()

	log.Print("The service is ready to listen and serve PORT: " + port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}