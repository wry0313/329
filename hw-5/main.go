package main

import (
	"fmt"
	"sync"
	"time"
)

var x int = 0
var max int = 1000

func increment(i int, m *TASLock, wg *sync.WaitGroup) {
	for {
		// wait for 2 sec
		time.Sleep(10 * time.Millisecond)
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

	var m TASLock

	for i := 0; i < max; i++ {
		wg.Add(1)
		go increment(i, &m, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s\n", elapsed) 
}
