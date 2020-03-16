package main

import (
	"log"
	"math/rand"
)

func main() {
	tree := initRedBlackTree()
	hashmap := make(map[int64]bool, 1000)
	for i := 0; i < 1000; i++ {
		hashmap[rand.Int63n(100000)] = true
	}
	for key, value := range hashmap {
		if value {
			tree.Add(key)
		}
	}
	log.Println()
}
