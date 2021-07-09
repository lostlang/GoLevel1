package main

import (
	"fmt"
	"sync"
)

func pow2(num int, wg *sync.WaitGroup) {
	out := num * num
	fmt.Println(out)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	arr := []int{2, 4, 6, 8, 10}
	wg.Add(len(arr))

	for _, i := range arr {
		go pow2(i, &wg)
	}

	wg.Wait()
}
