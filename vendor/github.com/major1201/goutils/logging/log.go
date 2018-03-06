package logging

import (
	"fmt"
	"time"
	"os"
)

const (
	Debug = iota
	Info
	Warning
	Error
	Fatal
)

const timeFormat = "2006-01-02T15:04:05"

type Logger struct {
	Name string
	IsUTC bool
	buf []byte
}

func New(name string) *Logger {
	logger := &Logger{}
	logger.Name = name
	logger.IsUTC = true
	return logger
}

func (logger *Logger) Log(level int, levelString, msg string) {
	t := time.Now()
	if logger.IsUTC == true {
		t = t.UTC()
	}
	logString := fmt.Sprintf("[%s %s %s] %s\n", t.Format(timeFormat), levelString, logger.Name, msg)
	logger.buf = logger.buf[:0]
	logger.buf = append(logger.buf, logString...)
	for _, w := range Writers {
		if level >= w.Level {
			w.Writer.Write(logger.buf)
		}
	}
}

func (logger *Logger) Debug(v ...interface{}) {
	logger.Log(Debug, "DEBUG", fmt.Sprint(v...))
}

func (logger *Logger) Debugf(format string, v ...interface{}) {
	logger.Log(Debug, "DEBUG", fmt.Sprintf(format, v...))
}

func (logger *Logger) Info(v ...interface{}) {
	logger.Log(Info, "INFO", fmt.Sprint(v...))
}

func (logger *Logger) Infof(format string, v ...interface{}) {
	logger.Log(Info, "INFO", fmt.Sprintf(format, v...))
}

func (logger *Logger) Warning(v ...interface{}) {
	logger.Log(Warning, "WARNING", fmt.Sprint(v...))
}

func (logger *Logger) Warningf(format string, v ...interface{}) {
	logger.Log(Warning, "WARNING", fmt.Sprintf(format, v...))
}

func (logger *Logger) Error(v ...interface{}) {
	logger.Log(Error, "ERROR", fmt.Sprint(v...))
}

func (logger *Logger) Errorf(format string, v ...interface{}) {
	logger.Log(Error, "ERROR", fmt.Sprintf(format, v...))
}

func (logger *Logger) Fatal(v ...interface{}) {
	logger.Log(Fatal, "FATAL", fmt.Sprint(v...))
	os.Exit(1)
}

func (logger *Logger) Fatalf(format string, v ...interface{}) {
	logger.Log(Fatal, "FATAL", fmt.Sprintf(format, v...))
	os.Exit(1)
}
