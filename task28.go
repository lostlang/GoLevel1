package main

import "fmt"

type intData struct {
	data int
}

type floatData struct {
	data float64
}

type adapter struct {
	data interface{}
}

func (a *adapter) printAdapt() {
	fmt.Println("adapted", a.data)
}

func main() {
	v := intData{12}
	v2 := floatData{0.2}

	var a = adapter{v}

	a.printAdapt()

	a = adapter{v2}

	a.printAdapt()
}
