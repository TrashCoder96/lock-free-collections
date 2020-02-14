package main

import "math/rand"

func main() {
	list := InitList()
	arr := make([]int, 1000, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = rand.Intn(10000)
	}
	for _, i := range arr {
		list.add(i)
	}
	for _, i := range arr {
		list.delete(i)
	}
}
