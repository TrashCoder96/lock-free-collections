package main

import (
	"log"
	"math/rand"
)

func main() {
	tree := initTree(5)
	array := make([]int64, 10000, 10000)
	for i := 0; i < 10000; i++ {
		array[i] = rand.Int63n(100000)
	}
	for _, item := range array {
		tree.Insert(item, "")
	}
	/*for _, item := range array {
		if item == 46637 {
			log.Println()
		}
		suc := tree.Delete(item)
		if !suc {
			log.Println()
		}
	}*/
	pointer := tree.root.internalNodeHead.childNode.internalNodeHead.childNode.internalNodeHead.childNode.internalNodeHead.childNode.leafHead
	for pointer != nil {
		pointer = pointer.nextKey
		log.Println(pointer.value)
	}
	log.Println(pointer)
}
