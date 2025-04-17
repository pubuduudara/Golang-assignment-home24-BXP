package logger

import (
	"log/slog"
	"os"
)

var (
	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
)

// Info logs an informational message
func Info(msg string) {
	logger.Info(msg)
}

// Warn logs a warning message
func Warn(msg string) {
	logger.Warn(msg)
}

// Error logs an error message
func Error(err error, msg ...string) {
	if err != nil {
		if len(msg) > 0 {
			logger.Error(msg[0], "error", err)
		} else {
			logger.Error("Error occurred", "error", err)
		}
	}
}
