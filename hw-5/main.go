package main

import (
	"fmt"
	"sync"
	"time"
)

var x int = 0

const numGoRoutine int = 200
const max = 20000

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

func incrementCLH(i int, m *CLHLock, wg *sync.WaitGroup) {
	for {

		myNode := &QNode{locked: false}
		myPred := &QNode{locked: false}
		m.Lock(i, myNode, myPred)
		if x >= max {
			m.Unlock(i, myNode, myPred)
			wg.Done()
			return
		}
		fmt.Println(x)
		x++
		m.Unlock(i, myNode, myPred)
	}
}

func main() {

	start := time.Now()
	var wg sync.WaitGroup

	// var m TASLock
	// var m TTASLock
	// minDelay := 1 * time.Millisecond
	// maxDelay := 10 * time.Millisecond
	// var m Lock = NewBackoffLock(minDelay, maxDelay)

	// var m = NewALock(max)
	var m = NewCLHLock()

	for i := 0; i < numGoRoutine; i++ {
		wg.Add(1)
		go incrementCLH(i, m, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s\n", elapsed)
}
