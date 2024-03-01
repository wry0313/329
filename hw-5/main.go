package main

import (
	"fmt"
	"sync"
	"time"
)

var x int = 0

const numGoRoutine int = 4
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

func incrementALock(i int, m *ALock, wg *sync.WaitGroup) {
	for {

		slot := m.Lock(i)
		if x >= max {
			m.Unlock(i, int(slot))
			wg.Done()
			return
		}
		fmt.Println(x)
		x++
		m.Unlock(i, int(slot))
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

func incrementMCS(i int, m *MCSLock, wg *sync.WaitGroup) {
	for {
	myNode := NewMCS_QNode()
		m.Lock(myNode)
		if x >= max {
			m.Unlock(myNode)
			wg.Done()
			return
		}
		fmt.Println(x)
		x++
		m.Unlock(myNode)
	}
}

func main() {

	start := time.Now()
	var wg sync.WaitGroup

	// var m TASLock
	// var m TTASLock

	// minDelay := 1 * time.Millisecond
	// maxDelay := 5 * time.Millisecond
	// var m Lock = NewBackoffLock(minDelay, maxDelay)

	// var m = NewALock(max)
	// var m = NewCLHLock()
	var m = NewMCSLock()

	for i := 0; i < numGoRoutine; i++ {
		wg.Add(1)
		// go increment(i, m, &wg)
		// go incrementALock(i, m, &wg)
		// go incrementCLH(i, m, &wg)
		go incrementMCS(i, m, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s\n", elapsed)
}
