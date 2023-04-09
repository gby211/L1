package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// конструктор Point
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) distance(q *Point) float64 {
	return math.Sqrt((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y))
}

func main() {
	i := NewPoint(0, 0)
	j := NewPoint(3, 4)
	ans := fmt.Sprintf("%f", i.distance(j))
	fmt.Println(ans) // 5
}
