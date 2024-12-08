package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x, y}
}

// Формула d = sqrt((x1 - x2)**2 + (y1 - y2)**2)
func (p1 *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func (p1 Point) String() string {
	return fmt.Sprintf("(%.3f, %.3f)", p1.x, p1.y)
}

func main() {
	p1 := &Point{3, 5}
	p2 := &Point{10, 9}

	fmt.Printf("Точка 1: %v\n", p1)
	fmt.Printf("Точка 2: %v\n", p2)
	fmt.Printf("Расстоянием между точками: %.3f\n", p1.Distance(p2))
}
