package main

import (
	"fmt"
	"sync"
)

func arrayPrinter(index *int, wg *sync.WaitGroup) {
	for {
		*index++
		fmt.Println(*index)
	}
}

func main() {
	var i = 0
	var wg sync.WaitGroup

	wg.Add(1)

	go arrayPrinter(&i, &wg)
	go arrayPrinter(&i, &wg)

	wg.Wait()
}
