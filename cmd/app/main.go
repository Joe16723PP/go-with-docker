package main

import (
	"fmt"
	"go-with-docker/app/router"
	"go-with-docker/config"
	lr "go-with-docker/util/logger"
	"net/http"
)

func main() {
	appConf := config.AppConfig()
	address := fmt.Sprintf(":%d", appConf.Server.Port)
	appRouter := router.New()
	logger := lr.New(appConf.Debug)

	logger.Info().Msgf("running on localhost %s", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Msg("Create server failure")
	}
}
