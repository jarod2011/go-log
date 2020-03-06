package go_log

import (
	"io"
	"log"
	"os"
)

var (
	level              Level
	de, in, wa, er, fa *mylog
	prefix             string
)

func init() {
	level = Info
	de = newMyLog(Debug, os.Stdout)
	in = newMyLog(Info, os.Stdout)
	wa = newMyLog(Warn, os.Stdout)
	er = newMyLog(Error, os.Stderr)
	fa = newMyLog(Fatal, os.Stderr)
}

type Level uint8

const (
	Debug Level = iota
	Info
	Warn
	Error
	Fatal
)

var names = [...]string{
	Debug: "DEBUG",
	Info:  "INFO",
	Warn:  "WARN",
	Error: "ERROR",
	Fatal: "FATAL",
}

func (l Level) String() string {
	if int(l) < len(names) {
		return names[l]
	}
	return ""
}

func (l Level) Prefix() string {
	name := l.String()
	if name != "" {
		return "[" + name + "]"
	}
	return ""
}

func SetLevel(l Level) {
	level = l
}

func GetLevel() Level {
	return level
}

func GetPrefix() string {
	return prefix
}

func SetPrefix(s string) {
	prefix = s
	de.SetPrefix(Debug.Prefix() + prefix)
	in.SetPrefix(Info.Prefix() + prefix)
	wa.SetPrefix(Warn.Prefix() + prefix)
	er.SetPrefix(Error.Prefix() + prefix)
	fa.SetPrefix(Fatal.Prefix() + prefix)
}

type Logger interface {
	Log(v ...interface{})
	Logf(format string, v ...interface{})
	SetWriter(writer io.Writer)
}

type mylog struct {
	l Level
	*log.Logger
}

func (m *mylog) Log(v ...interface{}) {
	if m.l >= level {
		if m.l == Fatal {
			m.Fatal(v...)
		} else {
			m.Print(v...)
		}
	}
}

func (m *mylog) Logf(format string, v ...interface{}) {
	if m.l >= level {
		if m.l == Fatal {
			m.Fatalf(format, v...)
		} else {
			m.Printf(format, v...)
		}
	}
}

func (m *mylog) SetWriter(writer io.Writer) {
	m.SetOutput(writer)
}

func newMyLog(level Level, writer io.Writer) *mylog {
	return &mylog{level, log.New(writer, level.Prefix(), log.LstdFlags)}
}

func D() Logger {
	return de
}

func I() Logger {
	return in
}

func W() Logger {
	return wa
}

func E() Logger {
	return er
}

func F() Logger {
	return fa
}
