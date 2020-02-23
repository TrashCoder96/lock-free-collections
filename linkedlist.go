package main

import "sync"

type node struct {
	Value   int
	Next    *node
	Mutex   sync.Mutex
	Removed bool
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
		for cursor != nil && cursor.Value < value {
			previousNode = cursor
			cursor = cursor.Next
		}
		previousNode.Mutex.Lock()
		success := false
		if previousNode.Next == cursor && !previousNode.Removed {
			newNode.Next = cursor
			previousNode.Next = newNode
			success = true
		}
		previousNode.Mutex.Unlock()
		if success {
			return
		}
	}
}

func (l *OrderedLinkedList) delete(value int) {
	for !l.deleteInCycle(value) {
	}
}

func (l *OrderedLinkedList) deleteInCycle(value int) bool {
	cursor := l.Head
	var previousNode *node
	for cursor != nil && cursor.Value != value {
		previousNode = cursor
		cursor = cursor.Next
	}
	if cursor == nil {
		panic("Key not found")
	}
	previousNode.Mutex.Lock()
	defer previousNode.Mutex.Unlock()
	cursor.Mutex.Lock()
	defer cursor.Mutex.Unlock()
	condition := previousNode.Next == cursor && !previousNode.Removed && !cursor.Removed
	if cursor.Next != nil {
		condition = condition && !cursor.Next.Removed
	}
	if condition {
		cursor.Removed = true
		previousNode.Next = cursor.Next
		return true
	}
	return false
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
