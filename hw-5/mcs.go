package main

import (
	"sync/atomic"
)

type MCS_QNode struct {
	locked bool
	next   *MCS_QNode
}

type MCSLock struct {
	tail *atomic.Pointer[MCS_QNode]
}

func NewMCS_QNode() *MCS_QNode {
	return &MCS_QNode{next: nil, locked: false}
}

func NewMCSLock() *MCSLock {

	var newAtomicPointer atomic.Pointer[MCS_QNode]
	newAtomicPointer.Store(nil)
	return &MCSLock{tail: &newAtomicPointer}
}

func (x *MCSLock) Lock(myNode *MCS_QNode) {

	pred := x.tail.Swap(myNode)
	if pred != nil {
		myNode.locked = true
		pred.next = myNode
		for myNode.locked {
		}
	}
}

func (x *MCSLock) Unlock(myNode *MCS_QNode) {
	if myNode.next == nil {
		if x.tail.CompareAndSwap(myNode, nil) {
			return
		}
		for myNode.next == nil {
		}
	}
	nextNode := myNode.next
	nextNode.locked = false
}
