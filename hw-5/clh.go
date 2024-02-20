package main

import (
	"sync/atomic"
)

type QNode struct {
	locked bool
}

type CLHLock struct {
	tail *atomic.Pointer[QNode] // Stores *QNode
	// nodes []*QNode
}

func NewCLHLock() *CLHLock {
	lock := &CLHLock{
		tail: &atomic.Pointer[QNode]{},
	}
	lock.tail.Store(&QNode{locked: false})
	return lock
}

func (cl *CLHLock) Lock(i int, myNode, myPred *QNode) {
	node := myNode
	node.locked = true
	pred := cl.tail.Swap(node)
	myPred = pred
	for pred.locked {
		// Busy wait
	}
}

func (cl *CLHLock) Unlock(i int, myNode, myPred *QNode) {
	qnode := myNode
	qnode.locked = false
	myNode = myPred
}
