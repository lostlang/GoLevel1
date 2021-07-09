package main

import "fmt"

func uniqueString(s string) bool {
	con := make(map[rune]int)

	for _, i := range s {
		if con[i] == 1 {
			return false
		} else {
			con[i] = 1
		}
	}
	return true
}

func main() {
	var s1 = "123456789"
	var s2 = "112255"

	fmt.Println(uniqueString(s1))
	fmt.Println(uniqueString(s2))
}
