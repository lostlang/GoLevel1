package main

import (
	"fmt"
	"sync"
)

func pow2(num int, wg *sync.WaitGroup, sum *int) {
	out := num * num
	*sum += out
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	arr := []int{2, 4, 6, 8, 10}
	sum_sq := 0
	wg.Add(len(arr))

	for _, i := range arr {
		go pow2(i, &wg, &sum_sq)
	}

	wg.Wait()
	fmt.Println(sum_sq)
}
