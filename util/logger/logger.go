package logger

import (
	"github.com/rs/zerolog"
	"io"
	"os"
)

type Logger struct {
	logger *zerolog.Logger
}

func New(isDebug bool) *Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &Logger{logger: &logger}
}
func NewConsole(isDebug bool) *Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &Logger{logger: &logger}
}

// Output duplicates the global logger and sets w as its output.
func (l *Logger) Output(w io.Writer) zerolog.Logger {
	return l.logger.Output(w)
}

// Info You must call Msg on the returned event in order to send the event.
func (l *Logger) Info() *zerolog.Event {
	return l.logger.Info()
}

func (l *Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal()
}
