package main

import (
	"log"
	"math/rand"
)

func main() {
	tree := initRedBlackTree()
	hashmap := make(map[int64]bool, 500000)
	for i := 0; i < 500000; i++ {
		u := rand.Int63n(9000000000000000000)
		hashmap[u] = true
	}
	for key, value := range hashmap {
		if value {
			tree.Add(key) //festinating
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
	for key, value := range hashmap {
		if value {
			b := tree.Delete(key)
			if !b {
				log.Println(key)
			}
		}
	}
}
