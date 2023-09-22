package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// Возращаем адресс стуктуры для глобального изменения и доступа
func new_point(x, y float64) *Point {
	return &Point{x, y}
}

// Считаем по формуле sqrt((x0 - x1)^2 + (y0 - y1)^2)
func distance(p1, p2 *Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := new_point(0, 0)
	p2 := new_point(3, 4)
	fmt.Printf("Расстояние %v\n", distance(p1, p2))
}
