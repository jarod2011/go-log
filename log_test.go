package go_log_test

import (
	log "github.com/jarod2011/go-log"
	"testing"
	"time"
)

var timeStart time.Time

func TestFlog_Log(t *testing.T) {
	log.SetLevel(log.Info)
	timeStart = time.Now()
	t1()
}

func t1() {
	for i := 0; i < 10; i++ {
		t2(i)
		time.Sleep(time.Millisecond * 5)
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
