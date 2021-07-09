package main

import (
	"fmt"
	"math"
)

type point2D struct {
	x, y float64
}

func NewPoint(x, y float64) Point2D {
	return point2D{x, y}
}

func getDistance(a, b point2D) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func main() {
	a := NewPoint(1, 2)
	b := NewPoint(10, -1)

	fmt.Println(GetDistance(a, b))
}
