package server

import (
	"github.com/jinzhu/gorm"
	"go-with-docker/util/logger"
)

type App struct {
	logger *logger.Logger
	db     *gorm.DB
}

func New(logger *logger.Logger, db *gorm.DB) *App {
	return &App{
		logger: logger,
		db:     db,
	}
}
func (app *App) Logger() *logger.Logger {
	return app.logger
}
