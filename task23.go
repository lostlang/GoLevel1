package main

import "fmt"

func binary(arr []int, num int) int {
	var k = len(arr) / 2
	for i := 0; i < len(arr); {
		switch {
		case arr[i] == num:
			{
				return i
			}
		case arr[i] < num:
			{
				i += k
			}
		case arr[i] > num:
			{
				i -= k
			}
		}
		if k > 1 {
			k = k / 2
		}
	}
	return -1
}

func main() {
	var arr = make([]int, 100)
	arr[99] = 20

	fmt.Println(binary(arr, 20))
}
