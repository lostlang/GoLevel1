package main

import (
	"fmt"
	"reflect"
)

func toArray(data interface{}) interface{} {
	switch reflect.TypeOf(data).String() {
	case "int":
		return intToArray(data.(int))
	case "string":
		return stringToArray(data.(string))
	default:
		{
			panic("error type")
		}
	}
}

func intToArray(num int) interface{} {
	var out = make([]int, 0)
	for num > 0 {
		out = append(out, num%10)
		num /= 10
	}
	return interface{}(out)
}

func stringToArray(s string) interface{} {
	var out = make([]rune, 0)
	for _, i := range s {
		out = append(out, i)
	}
	return interface{}(out)
}

func poli(arr interface{}) bool {
	switch reflect.TypeOf(arr).String() {
	case "[]int32":
		return poliRune(arr.([]rune))
	case "[]int":
		return poliInt(arr.([]int))
	default:
		{
			panic("error type")
		}
	}
}

func poliRune(arr []rune) bool {
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}
	return true
}

func poliInt(arr []int) bool {
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var i1 = 2515812
	var s1 = "oioork"
	var s2 = "12321"
	var i2 = 123321

	fmt.Println(poli(toArray(s1)))
	fmt.Println(poli(toArray(s2)))
	fmt.Println(poli(toArray(i1)))
	fmt.Println(poli(toArray(i2)))
}
