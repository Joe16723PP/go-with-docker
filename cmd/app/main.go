package main

import (
	"fmt"
	"go-with-docker/app/router"
	"go-with-docker/app/server"
	"go-with-docker/config"
	lr "go-with-docker/util/logger"
	"net/http"
)

func main() {
	appConf := config.AppConfig()
	address := fmt.Sprintf(":%d", appConf.Server.Port)
	logger := lr.New(appConf.Debug)
	application := server.New(logger)
	appRouter := router.New(application)
	logger.Info().Msgf("running on localhost %s", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Create server failure")
	}
}
