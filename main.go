package main

import (
	"log"
	"net/http"
	"os"
	"context"
	"syscall"
	"os/signal"

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
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	router := handlers.Router()
	srv := &http.Server {
		Addr: ":" + port,
		Handler: router,
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	log.Print("The service is ready to listen and serve PORT: " + port)

	killSignal := <-stop
		switch killSignal {
		case os.Interrupt:
			log.Print("Got SIGINT...")
		case syscall.SIGTERM:
			log.Print("Got SIGTERM...")
		}

		log.Print("The service is shutting down...")
		srv.Shutdown(context.Background())
		log.Print("Server gracefully stopped")
}