package main

import (
	"log"
	"math/rand"
)

func main() {
	tree := initRedBlackTree()
	hashmap := make(map[int64]bool, 50000)
	for i := 0; i < 50000; i++ {
		hashmap[rand.Int63n(10000000)] = true
	}
	for key, value := range hashmap {
		if value {
			tree.Add(key)
		}
	}
	for key, value := range hashmap {
		if value {
			b := tree.Find(key)
			if !b {
				log.Println(key)
			}
		}
	}
	log.Println()
}
