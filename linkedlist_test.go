package main

import (
	"math/rand"
	"sync"
	"testing"
)

func BenchmarkLinkedListBy100Goroutines(b *testing.B) {
	list := InitList()
	for m := 0; m < b.N; m++ {
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
	}
	b.Logf("Count of items : %d\n", list.count())
}

func BenchmarkLinkedListByOneGoroutine(b *testing.B) {
	list := InitList()
	for m := 0; m < b.N; m++ {
		for i := 0; i < 100; i++ {
			func(n int) {
				for j := n * 1000; j < (n+1)*1000; j++ {
					p := rand.Intn(100)
					list.add(p)
					if j%2 == 0 {
						list.delete(p)
					}
				}
			}(i)
		}
	}
	b.Logf("Count of items : %d\n", list.count())
}
