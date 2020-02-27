package main

import (
	"log"
	"math/rand"
)

func main() {
	tree := initTree(3)
	array := make([]int64, 2500, 2500)
	for i := 0; i < 2500; i++ {
		array[i] = rand.Int63n(100000)
	}
	for _, item := range array {
		tree.Insert(item, "")
	}
	for _, item := range array {
		s := tree.Delete(item)
		if !s {
			log.Println()
		}
	}
	log.Println()
}
