package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func newPoint(x, y float64) point {
	return point{
		x: x,
		y: y,
	}
}

func calcDistance(point1, point2 point) float64 {
	return math.Sqrt(math.Pow(point1.x-point2.x, 2) + math.Pow(point1.y-point2.y, 2))
}

func main() {
	point1 := newPoint(6.5, 3.8)
	point2 := newPoint(10.7, 14.2)
	fmt.Println(calcDistance(point1, point2))
}
