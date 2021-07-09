package main

import "fmt"

func arrToChan(arr []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, i := range arr {
			out <- i
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()

	return out
}

func main() {
	array := []int{10, 5, 11, 4, 13}
	ch := arrToChan(array)
	sqCh := sq(ch)

	fmt.Println(<-sqCh)
}
