package main

import (
	"fmt"
)

type Human struct {
	gender string
	name   string
}

type Actor struct {
	*Human
}

func main() {
	prot := Human{"male", "prot"}
	fmt.Println(prot)

	actor := Actor{&prot}
	actor.gender = "none"

	fmt.Println(actor)

	fmt.Println(prot)
}
