package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/Toffee-iZt/gomine/common"
	"github.com/fatih/color"
)

// Level ...
type Level = common.Bits

// Log levels
const (
	INFO Level = 1 << iota
	WARN
	FAIL
	DATA
)

//
const (
	BasicLog = INFO | WARN | FAIL
	FullLog  = BasicLog | DATA
)

// New creates new log object.
func New(source string, w io.Writer, level Level, color bool) *Logger {
	if w == nil {
		w = os.Stderr
	}
	return &Logger{source, w, level, color}
}

// Logger ...
type Logger struct {
	source string
	writer io.Writer
	level  Level
	color  bool
}

func (l *Logger) log(level Level, col func(string, ...interface{}) string, format string, msg ...interface{}) {
	if !l.level.Has(level) {
		return
	}

	h, m, s := time.Now().Clock()

	var lvl string

	switch level {
	case INFO:
		lvl = "INFO"
	case WARN:
		lvl = "WARN"
	case FAIL:
		lvl = "FAIL"
	case DATA:
		lvl = "DATA"
	}

	var name string
	var t string

	if l.color {
		name = color.WhiteString(l.source)
		lvl = col(lvl)
		t = color.GreenString("%d:%d:%d", h, m, s)
	} else {
		name = l.source
		t = fmt.Sprintf("%d:%d:%d", h, m, s)
	}

	fmt.Fprintf(l.writer, "[%s] [%s] [%s] %s\n", t, lvl, name, fmt.Sprintf(format, msg...))
}

// Info ...
func (l *Logger) Info(format string, msg ...interface{}) {
	l.log(INFO, color.CyanString, format, msg...)
}

// Warn ...
func (l *Logger) Warn(format string, msg ...interface{}) {
	l.log(WARN, color.YellowString, format, msg...)
}

// Fail ...
func (l *Logger) Fail(format string, msg ...interface{}) {
	l.log(FAIL, color.RedString, format, msg...)
}

// Data ...
func (l *Logger) Data(format string, msg ...interface{}) {
	l.log(DATA, color.MagentaString, format, msg...)
}
