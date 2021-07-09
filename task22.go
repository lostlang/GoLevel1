package main

import (
	"fmt"
)

func main() {
	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	Quicksort(ar)
	fmt.Println(ar)
}

func Quicksort(ar []int) {
	if len(ar) <= 1 {
		return
	}
	split := partition(ar)
	Quicksort(ar[:split])
	Quicksort(ar[split:])
}

func partition(ar []int) int {
	middle := ar[len(ar)/2]
	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < middle; left++ {
		}
		for ; ar[right] > middle; right-- {
		}
		if left >= right {
			return right
		}
		swap(ar, left, right)
	}
}

func swap(ar []int, i, j int) {
	ar[i] = ar[i] ^ ar[j]
	ar[j] = ar[i] ^ ar[j]
	ar[i] = ar[i] ^ ar[j]
}
