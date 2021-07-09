package main

import (
	"fmt"
	"time"
)

func timer(t int) {
	timer := time.After(time.Duration(t))
	<-timer
}

func main() {
	var second = 1000000000
	fmt.Println(1, time.Now())
	timer(second)
	fmt.Println(2, time.Now())
}
