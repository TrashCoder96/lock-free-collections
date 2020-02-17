package main

import (
	"log"
	"sync"
)

func main() {
	list := InitList()
	var wg sync.WaitGroup
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func(n int) {
			for j := n * 1000; j < (n+1)*1000; j++ {
				list.add(7)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Println(list.count())
	log.Println("")
}
