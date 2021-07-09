package main

import "fmt"

func arrayToCounter(arr []string) map[string]int {
	out := make(map[string]int)

	for _, i := range arr {
		out[i] += 1
	}

	return out
}

func main() {
	in := []string{"cat", "cat", "dog", "cat", "tree"}
	out := arrayToCounter(in)

	fmt.Println(out)
}
