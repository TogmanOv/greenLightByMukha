package main

import (
	"time"
)

// Extra test pass
func ExecutePipeline(jobs ...job) {

	inCh := make(chan interface{}, 3)
	var outCh chan interface{}

	for i := 0; i < len(jobs); i++ {

		outCh = execTask(jobs[i], inCh)
		inCh = make(chan interface{}, 3)

		timeout := time.Duration(i*100+7) * time.Millisecond
		if i > 1 {
			timeout = time.Duration(1) * time.Millisecond
		}

	LOOP:
		for {
			select {
			case v := <-outCh:
				inCh <- v
			case <-time.After(timeout):
				close(outCh)
				break LOOP
			}
		}
	}
}

func execTask(j job, inpCh chan interface{}) chan interface{} {
	outCh := make(chan interface{}, 1)
	go j(inpCh, outCh)
	return outCh
}
