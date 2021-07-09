package main

import (
	"fmt"
	"sync"
)

func reader(arr [][]int, out chan []int) {
	for _, i := range arr {
		out <- i
	}
	close(out)
}

func adder(in <-chan []int, out *map[int]int, mu *sync.Mutex) {
	for val := range in {
		mu.Lock()
		myMap := *out
		myMap[val[0]] = val[1]
		mu.Unlock()
	}
}

func main() {
	var in [][]int
	var out = make(map[int]int)
	var conn = make(chan []int)
	var mu sync.Mutex
	var adders = 2

	for i := 0; i < 10; i++ {
		in = append(in, []int{i, i * i})
	}

	go reader(in, conn)

	for i := 0; i < adders; i++ {
		go adder(conn, &out, &mu)
	}

	fmt.Println(out)
}
