package tlog

import (
	"fmt"
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

type LogStruct struct {
	Callfunc string
	Message  string
}

func NewLogStruct(svc, sms string) *LogStruct {
	return &LogStruct{
		Callfunc: svc,
		Message:  sms,
	}
}

func (l *LogStruct) toString() string {
	return fmt.Sprintf("{Service: %s, Message: %s}", l.Callfunc, l.Message)
}

func Debug(svc, msg string, keyvals ...interface{}) {
	logger.Debug(NewLogStruct(svc, msg).toString(), keyvals...)
}

func Info(svc, msg string, keyvals ...interface{}) {
	logger.Info(NewLogStruct(svc, msg).toString(), keyvals...)
}

func Warn(svc, msg string, keyvals ...interface{}) {
	logger.Warn(NewLogStruct(svc, msg).toString(), keyvals...)
}

func Error(svc, msg string, keyvals ...interface{}) {
	logger.Error(NewLogStruct(svc, msg).toString(), keyvals...)
}
