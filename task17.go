package main

import (
	"fmt"
	"reflect"
)

type testInterface struct {
	data interface{}
}

func (i testInterface) typeOf() interface{} {
	return reflect.TypeOf(i.data)
}

func main() {
	var s1 = testInterface{"string"}
	var i1 = testInterface{1}
	var f1 = testInterface{1.1}

	fmt.Println(s1, s1.typeOf())
	fmt.Println(i1, i1.typeOf())
	fmt.Println(f1, f1.typeOf())
}
