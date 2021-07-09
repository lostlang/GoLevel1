package main

import (
	"fmt"
	"time"
)

func reader(arr []int, out chan int) {
	for _, i := range arr {
		out <- i
	}
}

func writer(in chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	var arr = make([]int, 100)
	var c = make(chan int)

	go reader(arr, c)
	go writer(c)

	time.Sleep(time.Second * 1)
}
