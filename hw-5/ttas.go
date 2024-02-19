package main

import (
	"sync/atomic"
)


type TTASLock struct {
	lock atomic.Bool
}

func (m *TTASLock) Lock(i int) {
	for {
		if !m.lock.Load() {
			if m.lock.CompareAndSwap(false, true) {
				return
			}
		}
	}
}

func (m *TTASLock) Unlock(i int) {
	m.lock.Store(false)
}
