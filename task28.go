package main

import (
	"fmt"
)

func reverse(s string) string {
	var words = make([]string, 0)
	var word, out string

	for _, i := range s {
		if i != ' ' {
			word += string(i)
		} else {
			words = append(words, word)
			word = ""
		}
	}
	words = append(words, word)

	for i := len(words) - 1; i >= 0; i-- {
		out += words[i] + " "
	}
	return out
}

func main() {
	var s1 = "set Тест string Может быть"

	fmt.Println(s1)
	fmt.Println(reverse(s1))
}
