package log

import (
	"log"
	"runtime"
)

const logLevel = -1

type Log struct {
	level int
	name  string
	color string
}

var (
	LOG   = Log{level: 0, name: "LOG", color: "\033[0;32m"}
	WARN  = Log{level: 5, name: "WARN", color: "\033[0;33m"}
	ERROR = Log{level: 7, name: "ERROR", color: "\033[0;31m"}
)

func (l *Log) Message(v string) *string {
	_, file, line, _ := runtime.Caller(1)
	if l.level > logLevel {
		log.Printf("%s[%s] %s:%d %s\033[0m\n", l.color, l.name, file, line, v)
	}
	return &v
}
