package main

import (
	"sync/atomic"
	"unsafe"
)

type node struct {
	Value int
	Next  *node
}

//OrderedLinkedList struct
type OrderedLinkedList struct {
	Head *node
}

//InitList struct
func InitList() *OrderedLinkedList {
	return &OrderedLinkedList{
		Head: &node{
			Value: -1,
		},
	}
}

func (l *OrderedLinkedList) add(value int) {
	newNode := &node{
		Value: value,
	}
	for {
		cursor := l.Head
		var previousNode *node
		var oldValue unsafe.Pointer
		for cursor != nil && cursor.Value < value {
			previousNode = cursor
			oldValue = unsafe.Pointer(previousNode.Next)
			cursor = cursor.Next
		}
		newNode.Next = previousNode.Next
		newValue := unsafe.Pointer(newNode)
		addr := (*unsafe.Pointer)(unsafe.Pointer(&previousNode.Next))
		if atomic.CompareAndSwapPointer(addr, oldValue, newValue) {
			return
		}
	}
}

func (l *OrderedLinkedList) delete(value int) {
	cursor := l.Head
	var previousNode *node
	for cursor != nil && cursor.Value != value {
		previousNode = cursor
		cursor = cursor.Next
	}
	if cursor == nil {
		panic("Key not found")
	}
	previousNode.Next = previousNode.Next.Next
}

func (l *OrderedLinkedList) count() int {
	count := 0
	cursor := l.Head
	for cursor != nil {
		cursor = cursor.Next
		count++
	}
	return count - 1
}
