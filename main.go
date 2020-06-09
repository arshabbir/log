package main

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"

	log "github.com/sirupsen/logrus"
)

var gid int64

var Log *log.Logger //Intentionally making it an public variable

func init() {

	Log = &log.Logger{
		Out:       os.Stdout,
		Formatter: &log.JSONFormatter{},
		Level:     log.DebugLevel,
	}

}

func main() {

	var wg sync.WaitGroup
	for i := 1; i < 100; i++ {
		wg.Add(i)
		grid := atomic.AddInt64(&gid, 1)
		go r1(grid, &wg)
	}

	wg.Wait()

}

func r1(gid int64, wg *sync.WaitGroup) {

	var j = 0

	rid := fmt.Sprintf("%d", gid)

	for {

		j++

		Log.WithFields(log.Fields{
			"gid": rid,
		}).Info(j)

	}

	wg.Done() //This is not required & just a place holder

}
