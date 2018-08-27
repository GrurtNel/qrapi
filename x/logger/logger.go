package logger

import (
	"fmt"

	"github.com/golang/glog"
)

type Logger struct {
	Tag string
}

func NewLogger(tag string) *Logger {
	return &Logger{
		Tag: "[" + tag + "] ",
	}
}

func (l *Logger) Infof(format string, args ...interface{}) {
	format = l.Tag + format
	format = fmt.Sprintf(format, args...)
	glog.Infof(format)
}

func (l *Logger) Error(args ...interface{}) {
	glog.Error(fmt.Sprint(l.Tag, args))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	format = l.Tag + format
	format = fmt.Sprintf(format, args...)
	glog.Errorf(format)
}
