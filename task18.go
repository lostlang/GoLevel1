package main

import (
	"fmt"
)

func someAction(v []int8, b int8) {
	arr := make([]int8, len(v))
	copy(arr, v)
	arr[0] = 100
	arr = append(arr, b)
	fmt.Print(arr)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
