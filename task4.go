package main

import (
	"fmt"
)

func worker(idWorker int, input <-chan int, dead <-chan bool) {
	for {
		select {
		case <-dead:
			return
		case val := <-input:
			fmt.Println(idWorker, val)
		}
	}
}

func deadAll(count int, dead chan bool) {
	for i := 0; i < count; i++ {
		dead <- true
	}
}

func main() {

	workers := 5
	inputs := make(chan int)
	dead := make(chan bool)

	for i := 0; i < workers; i++ {
		go worker(i, inputs, dead)
	}

	for i := 0; ; i++ {
		if i == 25000 {
			deadAll(workers, dead)
			close(inputs)
			break
		}
		inputs <- i
	}
}
