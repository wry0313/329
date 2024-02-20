package main

import (
	"fmt"
	"sync"
	"time"
)

var x int = 0
var max int = 1000

type Lock interface {
	Lock(i int)
	Unlock(i int)
}

func increment(i int, m Lock, wg *sync.WaitGroup) {
	for {
		m.Lock(i)
		if x >= max {
			m.Unlock(i)
			wg.Done()
			return
		}
		fmt.Println(x)
		x++
		m.Unlock(i)
	}
}

func main() {

	start := time.Now()
	var wg sync.WaitGroup

	// minDelay := 1 * time.Millisecond
	// maxDelay := 10 * time.Millisecond
	// var m Lock = NewBackoffLock(minDelay, maxDelay)

	// var m TASLock

	var m Lock = NewALock(max)

	for i := 0; i < max; i++ {
		wg.Add(1)
		go increment(i, m, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s\n", elapsed)
}
