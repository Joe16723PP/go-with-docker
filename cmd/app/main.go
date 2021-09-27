package main

import (
	"fmt"
	"go-with-docker/app/router"
	"go-with-docker/config"
	"log"
	"net/http"
)

func main() {
	appConf := config.AppConfig()
	address := fmt.Sprintf(":%d", appConf.Server.Port)
	appRouter := router.New()

	log.Printf("running on localhost %s", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
