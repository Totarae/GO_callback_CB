package main

import (
	"log"
	"net/http"
	"project/service1/handler"
)

func main() {
	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/callback", handler.CallbackHandler)

	// Create a server with timeouts for better performance
	server := &http.Server{
		Addr:         ":9090",
		Handler:      nil, // DefaultServeMux is used
		ReadTimeout:  5 * 60,
		WriteTimeout: 10 * 60,
		IdleTimeout:  15 * 60,
	}

	log.Printf("Starting server on :9090")
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Error starting server:", err)
	}
}
