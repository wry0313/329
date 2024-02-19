package main

import (
	"sync/atomic"
)


type TASLock struct {
	lock atomic.Bool
}

func (m *TASLock) Lock(i int) {
	// fmt.Println("Locking", i)
	for {
		if !m.lock.Load() {
			if m.lock.CompareAndSwap(false, true) {
				return
			}
		}
	}
}

func (m *TASLock) Unlock(i int) {
	// fmt.Println("Unlocking", i)
	m.lock.Store(false)
}
