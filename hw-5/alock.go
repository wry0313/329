package main

import (
	"sync/atomic"
)

type ALock struct {
	mySlotIndex []int
	flag        []bool
	size        int
	tail        atomic.Int32
}

func NewALock(capacity int) *ALock {
	ret := &ALock{
		flag:        make([]bool, capacity),
		size:        capacity,
		tail:        atomic.Int32{},
		mySlotIndex: make([]int, capacity),
	}
	ret.flag[0] = true
	return ret
}

func (l *ALock) Lock(i int) {
	slot := (l.tail.Add(1)-1) % int32(l.size)
	l.mySlotIndex[i] = int(slot)
	for !l.flag[slot] {
	}
}

func (l *ALock) Unlock(i int) {
	slot := l.mySlotIndex[i]
	l.flag[slot] = false
	l.flag[(slot+1)%l.size] = true
}
