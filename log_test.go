package go_log_test

import (
	"bytes"
	log "github.com/jarod2011/go-log"
	"sync/atomic"
	"testing"
	"time"
)

var timeStart time.Time
var count uint32

func TestFlog_Log(t *testing.T) {
	log.SetLevel(log.Info)
	log.SetFatalExit(false)
	timeStart = time.Now()
	t1()
	t.Log(count)
	if count != 8 {
		t.Errorf("count should equal 8 but %d not", count)
	}
	t.Logf("all used %v", time.Since(timeStart))
}

func TestLogs(t *testing.T) {
	buf := new(bytes.Buffer)
	log.SetPrefix("[test]")
	if log.GetPrefix() != "[test]" {
		t.Error("set log prefix failure")
	}
	log.SetFatalExit(false)
	log.SetLevel(log.Debug)
	if log.GetLevel() != log.Debug {
		t.Error("set log level failure")
	}
	logs := []log.Logger{log.D(), log.E(), log.W(), log.I()}
	for _, l := range logs {
		if buf.Len() > 0 {
			t.Error("buffer should len 0")
		}
		l.SetWriter(buf)
		l.Log(1)
		t.Log(buf.Len())
		if buf.Len() == 0 {
			t.Errorf("%v write to buffer failure", l)
		}
		buf.Reset()
		l.Logf("write %v", time.Now())
		t.Log(buf.Len())
		if buf.Len() == 0 {
			t.Errorf("%v write to buffer failure", l)
		}
		buf.Reset()
	}
	f(t)
}

func f(t *testing.T) {
	buf := new(bytes.Buffer)
	defer func() {
		err := recover()
		t.Log(err)
		if buf.Len() == 0 {
			t.Error("fatal write to buffer failure")
		}
	}()
	log.F().SetWriter(buf)
	log.F().Log("1")
}

func t1() {
	defer func() {
		if err := recover(); err != nil {
			log.E().Logf("recover %v", err)
		}
	}()
	for i := 0; i < 10; i++ {
		t2(i)
		time.Sleep(time.Millisecond * 5)
		atomic.AddUint32(&count, 1)
	}
}

func t3(index int) {
	for i := 0; i < index; i++ {
		<-time.After(time.Duration(index))
		log.D().Logf("index %d timeout now used %v", i, time.Since(timeStart))
	}
}

func t2(index int) {
	t3(index)
	if index == 8 {
		log.F().Logf("break at %d", index)
	}
	log.D().Logf("index %d", index)
}

func BenchmarkLog(b *testing.B) {
	log.SetLevel(log.Info)
	for i := 0; i < b.N; i ++ {
		log.D().Log(i)
	}
}