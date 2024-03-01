package main

import (
	"sync/atomic"
)

type TASLock struct {
	lock atomic.Bool
}

func (m *TASLock) Lock(i int) {
	for {
		if m.lock.CompareAndSwap(false, true) {
			return
		}
	}
}

func (m *TASLock) Unlock(i int) {
	m.lock.Store(false)
}

func (m *TASLock) IsLocked() bool {
	return m.lock.Load()
}


