package main

import "fmt"

func main() {
	var a, b int

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Print(a, b)
}
