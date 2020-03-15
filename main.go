package main

import (
	"log"
	"math/rand"
)

func main() {
	tree := initTree(3)
	hmap := make(map[int64]bool)
	for i := 0; i < 55000; i++ {
		hmap[rand.Int63n(10000000)] = true
	}
	log.Println(len(hmap))
	for key, value := range hmap {
		if value {
			if err := tree.Insert(key, ""); err != nil {
				log.Println(err)
			}
		}
	}
	for key, value := range hmap {
		if value {
			suc := tree.Delete(key)
			if !suc {
				log.Println("delete error")
			}
		}
	}
	log.Println()
}
