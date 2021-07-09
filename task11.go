package main

import (
	"fmt"
)

func arrayToCounter(array []int) map[int]int {
	out := make(map[int]int)

	for _, i := range array {
		out[i] += 1
	}

	return out
}

func intersection(c1, c2 map[int]int) map[int]int {
	out := make(map[int]int)

	for i := range c1 {
		if c1[i] > c2[i] {
			out[i] = c2[i]
		} else {
			out[i] = c1[i]
		}
	}
	for i := range c2 {
		if c1[i] > c2[i] {
			out[i] = c2[i]
		} else {
			out[i] = c1[i]
		}
	}

	return out
}

func main() {
	array1 := []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 4, 5, 7}
	array2 := []int{2, 3, 3, 7, 10, 5, 2}

	counter1 := arrayToCounter(array1)
	counter2 := arrayToCounter(array2)

	out := intersection(counter1, counter2)

	fmt.Println(out)

}
