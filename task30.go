package main

import (
	"fmt"
)

func main() {
	var p = make([]int, 20)
	var p2 = make([]int, len(p)-1)
	var i = 10

	for i := 0; i < 20; i++ {
		p[i] = i
	}

	copy(p2[:i], p[:i])
	copy(p2[i:], p[i+1:])

	fmt.Println(p2)
}
