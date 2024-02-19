package main

import (
	"math/rand"
	"sync/atomic"
	"time"
)

type Backoff struct {
	minDelay time.Duration
	maxDelay time.Duration
	limit    time.Duration
}

func NewBackoff(min, max time.Duration) *Backoff {
	return &Backoff{
		minDelay: min,
		maxDelay: max,
		limit:    min,
	}
}

func (b *Backoff) backoff() {
	delay := time.Duration(rand.Int63n(int64(b.limit)))
	time.Sleep(delay)
	if newLimit := 2 * b.limit; newLimit <= b.maxDelay {
		b.limit = newLimit
	} else {
		b.limit = b.maxDelay
	}
}

type BackoffLock struct {
	lock     atomic.Bool
	minDelay time.Duration
	maxDelay time.Duration
}

func NewBackoffLock(minDelay, maxDelay time.Duration) *BackoffLock {
	return &BackoffLock{
		lock:     atomic.Bool{},
		minDelay: minDelay,
		maxDelay: maxDelay,
	}
}

func (l *BackoffLock) Lock(i int) {
	backoff := NewBackoff(l.minDelay, l.maxDelay)
	for {
		if !l.lock.Load() {
			if l.lock.CompareAndSwap(false, true) {
				return
			}
		}
		backoff.backoff()
	}
}

func (l *BackoffLock) Unlock(i int) {
	l.lock.Store(false)
}