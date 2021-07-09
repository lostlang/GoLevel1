package main

import "fmt"

func reverse(s string) string {
	var runs = []rune(s)
	var out string

	for i := len(runs) - 1; i >= 0; i-- {
		out = out + string(runs[i])
	}

	return out
}

func main() {
	var s1 = "set Тест string Может быть"

	fmt.Println(s1)
	fmt.Println(reverse(s1))
}
