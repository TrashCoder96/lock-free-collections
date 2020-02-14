package main

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
	cursor := l.Head
	var previousNode *node
	for cursor != nil && cursor.Value <= value {
		previousNode = cursor
		cursor = cursor.Next
	}
	newNode := node{
		Value: value,
		Next:  cursor,
	}
	if previousNode != nil {
		previousNode.Next = &newNode
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
