package main

import "fmt"

func main() {
	tmp := []float64{1, 2, 4, 7, 7, 10, 11, 15, -10, -15, -19, -20}
	out := make(map[int][]float64)

	for _, i := range tmp {
		con := int(i) / 10 * 10
		out[con] = append(out[con], i)
	}

	fmt.Println(out)
}
