package go_log

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"
)

var (
	level          Level
	de, in, wa, er *mylog
	fa             *flog
	prefix         string
)

func init() {
	level = Info
	de = newMyLog(Debug, os.Stdout)
	in = newMyLog(Info, os.Stdout)
	wa = newMyLog(Warn, os.Stdout)
	er = newMyLog(Error, os.Stderr)
	fa = newFLog(os.Stderr)
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
		m.Print(v...)
	}
}

func (m *mylog) Logf(format string, v ...interface{}) {
	if m.l >= level {
		m.Printf(format, v...)
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

type flog struct {
	*log.Logger
}

func (l *flog) printStack() {
	buf := make([]byte, 1<<30)
	n := runtime.Stack(buf, true)
	if n > 0 {
		out := new(bytes.Buffer)
		fmt.Fprint(out, "==== Stack Start ====\n")
		fmt.Fprintf(out, "Time: %v\n", time.Now())
		fmt.Fprint(out, string(buf[:n]))
		fmt.Fprint(out, "==== Stack End ====\n")
		out.WriteTo(l.Logger.Writer())
	}
}

func (l *flog) Log(v ...interface{}) {
	l.printStack()
	l.Logger.Fatal(v...)
}

func (l *flog) Logf(format string, v ...interface{}) {
	l.printStack()
	l.Logger.Fatalf(format, v...)
}

func (f *flog) SetWriter(writer io.Writer) {
	f.Logger.SetOutput(writer)
}

func newFLog(writer io.Writer) *flog {
	return &flog{log.New(writer, Fatal.Prefix(), log.LstdFlags)}
}
