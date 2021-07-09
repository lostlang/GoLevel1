package main

import (
	"fmt"
)

func toInt(s string) int {
	var out int
	for _, i := range s {
		switch i {
		case '0':
			out = out*10 + 0
		case '1':
			out = out*10 + 1
		case '2':
			out = out*10 + 2
		case '3':
			out = out*10 + 3
		case '4':
			out = out*10 + 4
		case '5':
			out = out*10 + 5
		case '6':
			out = out*10 + 6
		case '7':
			out = out*10 + 7
		case '8':
			out = out*10 + 8
		case '9':
			out = out*10 + 9
		default:
			{
			}
		}
	}
	return out
}

func main() {
	var s1 = "251705518841   521747"

	fmt.Println(s1)
	fmt.Println(toInt(s1))
}
