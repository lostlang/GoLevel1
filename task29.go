package main

import (
	"fmt"
	"math/big"
)

func Mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}

func Div(x, y *big.Int) *big.Int {
	return big.NewInt(0).Div(x, y)
}

func Add(x, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}

func Sub(x, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}

func main() {
	var a = big.NewInt(1 << 62)
	var b = big.NewInt(1 << 61)

	fmt.Println(Mul(a, b))
	fmt.Println(Div(a, b))
	fmt.Println(Add(a, b))
	fmt.Println(Sub(a, b))
}
