package main

import (
	"fmt"
	"go-with-docker/config"
	"log"
	"net/http"
)

func main() {
	appConfig := config.AppConfig()

	address := fmt.Sprintf(":%d", appConfig.Server.Port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)
	log.Println("Starting server :8080")
	s := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  appConfig.Server.TimeoutRead,
		WriteTimeout: appConfig.Server.TimeoutWrite,
		IdleTimeout:  appConfig.Server.TimeoutIdle,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
