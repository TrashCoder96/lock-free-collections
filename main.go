package main

import (
	"log"
	"math/rand"
	"sync"
)

func main() {
	list := InitList()
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(n int) {
			for j := n * 1000; j < (n+1)*1000; j++ {
				p := rand.Intn(100)
				list.add(p)
				if j%2 == 0 {
					list.delete(p)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Printf("Count of items : %d\n", list.count())
}
