package server

import (
	"github.com/jinzhu/gorm"
	"go-with-docker/util/logger"
)

type App struct {
	logger *logger.Logger
	db     *gorm.DB
}

const (
	appErrDataAccessFailure   = "data access failure"
	appErrJsonCreationFailure = "json creation failure"
	appErrDataCreationFailure = "data creation failure"
	appErrFormDecodingFailure = "form decoding failure"
	appErrDataUpdateFailure   = "data update failure"
)

func New(logger *logger.Logger, db *gorm.DB) *App {
	return &App{
		logger: logger,
		db:     db,
	}
}
func (app *App) Logger() *logger.Logger {
	return app.logger
}
