package main

import (
	"math/rand"
	"testing"
)

func TestCorrectntessRedTree_ok(t *testing.T) {
	tree := initRedBlackTree()
	hashmap := make(map[int64]bool, 500000)
	for i := 0; i < 500000; i++ {
		hashmap[rand.Int63n(9000000000000000000)] = true
	}
	for key, value := range hashmap {
		if value {
			tree.Add(key)
		}
	}
	j := 0
	for key, value := range hashmap {
		if value {
			tree.Delete(key)
			j++
			if j == 250000 {
				break
			}
		}
	}
	i := 0
	processNode(tree.head, &i, t)
	if i != 250000 {
		t.FailNow()
	}
}

func processNode(node *redBlackNode, n *int, t *testing.T) {
	*n++
	if node.leftNode != nil {
		if !(node.leftNode.value < node.value) {
			t.FailNow()
		}
		if node.leftNode.parent != node {
			t.FailNow()
		}
		processNode(node.leftNode, n, t)
	}
	if node.rightNode != nil {
		if !(node.rightNode.value > node.value) {
			t.FailNow()
		}
		if node.rightNode.parent != node {
			t.FailNow()
		}
		processNode(node.rightNode, n, t)
	}
}
