package main

import (
	"sync/atomic"
)

type MCS_QNode struct {
	locked *atomic.Bool
	next   *atomic.Value
}

type MCSLock struct {
	tail *atomic.Pointer[MCS_QNode]
}

func NewMCS_QNode() *MCS_QNode {
	var new_atomic atomic.Value
	return &MCS_QNode{next: &new_atomic}
}

func NewMCSLock() *MCSLock {
	//newNode := NewMCS_QNode()

	var newAtomicPointer atomic.Pointer[MCS_QNode]
	newAtomicPointer.Store(nil)
	return &MCSLock{tail: &newAtomicPointer}
}

func (x *MCSLock) Lock(myNode *MCS_QNode) {
	qnode := *myNode

	pred := x.tail.Swap(&qnode)
	if pred != nil {
		qnode.locked.Store(true)
		pred.next.Store(&qnode)
		for qnode.locked.Load() {

		}
	}
}

func (x *MCSLock) Unlock(myNode *MCS_QNode) {
	qnode := *myNode

	if qnode.next.Load() == nil {

		if x.tail.CompareAndSwap(&qnode, nil) {
			return
		}

		for qnode.next.Load() == nil {
		}
	}
	nextNode := qnode.next.Load().(*MCS_QNode)
	nextNode.locked.Store(false)
}