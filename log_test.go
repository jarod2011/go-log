package go_log_test

import (
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