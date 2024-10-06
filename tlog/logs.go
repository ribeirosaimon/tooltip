package tlog

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func init() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	logger = slog.New(handler)
}

func Debug(message string, keyvals ...interface{}) {
	logger.Debug(message, keyvals...)
}
func Info(message string, keyvals ...interface{}) {
	logger.Info(message, keyvals...)
}

func Warn(message string, keyvals ...interface{}) {
	logger.Warn(message, keyvals...)
}

func Error(message string, keyvals ...interface{}) {
	logger.Error(message, keyvals...)
}
