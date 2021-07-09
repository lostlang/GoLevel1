package main

import (
	"fmt"
	"strings"
)

var JustString string

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		v += "A"
	}
	return v
}

func someFunc(toCreate int, needBack int) string {
	v := createHugeString(toCreate)
	var out strings.Builder
	for i, r := range v {
		if i < needBack {
			out.WriteRune(r)
		}
	}
	return out.String()
}

func main() {
	var create = 1 << 10
	var need = 100
	JustString = someFunc(create, need)
	fmt.Println(JustString)
}
