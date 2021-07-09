package main

import (
	"fmt"
)

func reverse(s string) []string {
	var words = make([]string, 0)
	var word string

	for _, i := range s {
		if i != ' ' {
			word += string(i)
		} else {
			words = append(words, word)
			word = ""
		}
	}
	words = append(words, word)

	for i := 0; i < len(words)/2; i++ {
		word = words[i]
		words[i] = words[len(words)-i-1]
		words[len(words)-i-1] = word
	}

	return words
}

func main() {
	var s1 = "set Тест string Может быть"

	fmt.Println(s1)
	fmt.Println(reverse(s1))
}
