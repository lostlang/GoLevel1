package main

import (
	"fmt"
)

type Human struct {
	gender string
	name   string
}

type Actor struct {
	Human
}

func main() {

	var actor Actor
	actor.gender = "none"

	fmt.Println(actor)

}
