package main

import (
	"fmt"
)

func bitSwap(num int, bit int, needBit int) (out int) {
	fmt.Printf("%b\n", num)
	switch needBit {
	case 0:
		out = num &^ (1 << bit)
	case 1:
		out = num | (1 << bit)
	}
	fmt.Printf("%b\n", out)
	return
}

func main() {
	var test_int = 2

	test_int = bitSwap(test_int, 0, 0)

	fmt.Println(test_int)
}
