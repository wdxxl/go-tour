package main

import (
	"fmt"
	"runtime/pprof"
	"time"
)

func potentiallyTimeConsumingCall(result chan struct{}) {
	<-time.After(10 * time.Millisecond)
	result <- struct{}{}
}

func runALotOfTasks() int {
	for i := 0; i < 100; i++ {
		r := make(chan struct{})
		go potentiallyTimeConsumingCall(r)
		select {
		case <-r:
		case <-time.After(5 * time.Millisecond):
			fmt.Println(pprof.Lookup("goroutine").Count())
			return pprof.Lookup("goroutine").Count()
		}
	}
	return pprof.Lookup("goroutine").Count()
}
