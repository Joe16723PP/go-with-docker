package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"go-with-docker/util/logger"
)

type App struct {
	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate
}

const (
	appErrDataAccessFailure   = "data access failure"
	appErrJsonCreationFailure = "json creation failure"
	appErrDataCreationFailure = "data creation failure"
	appErrFormDecodingFailure = "form decoding failure"
	appErrDataUpdateFailure   = "data update failure"
)

func New(logger *logger.Logger, db *gorm.DB, validator *validator.Validate) *App {
	return &App{
		logger:    logger,
		db:        db,
		validator: validator,
	}
}
func (app *App) Logger() *logger.Logger {
	return app.logger
}
