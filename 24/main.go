package main

import (
	"fmt"
	"math"
)

// point struct for point's coordinates
type point struct {
	x float64
	y float64
}

// newPoint create new point
func newPoint(x, y float64) point {
	return point{
		x: x,
		y: y,
	}
}

// calcDistance calculate distance using the pythagorean theorem
func calcDistance(point1, point2 point) float64 {
	return math.Sqrt(math.Pow(point1.x-point2.x, 2) + math.Pow(point1.y-point2.y, 2))
}

func main() {
	//create two points
	point1 := newPoint(6.5, 3.8)
	point2 := newPoint(10.7, 14.2)
	//calculate and print distance
	fmt.Println(calcDistance(point1, point2))
}
