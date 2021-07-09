package main

import (
	"fmt"
	"sync"
)

func arrayPrinter(arr *[]int, index *int, wg *sync.WaitGroup) {
	var array = *arr
	defer func() {
		wg.Done()
	}()
	for {
		if *index < len(*arr)-1 {
			*index++
			fmt.Println(*index, array[*index])
		} else {
			break
		}
	}
}

func main() {
	var arr = make([]int, 100)
	var i = 0
	var wg sync.WaitGroup

	wg.Add(2)

	go arrayPrinter(&arr, &i, &wg)
	go arrayPrinter(&arr, &i, &wg)

	wg.Wait()
}
