package main

import (
	"fmt"
	"math"
)

// Стурктура точки
type Point struct {
	x, y float64
}

// Функция создания точки
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Метод получения дистанции до другой точки
func (p Point) GetDistant(another Point) float64 {
	zx := math.Pow((p.x - another.x), 2)
	zy := math.Pow((p.y - another.y), 2)
	return math.Sqrt(zx + zy)
}

func main() {
	// Нулевая точки
	p1 := NewPoint(0, 0)
	// Счет дистацнии до массива точке
	for i := -3; i < 4; i++ {
		for j := -3; j < 4; j++ {
			p2 := NewPoint(float64(i), float64(j))
			fmt.Println(p1.GetDistant(*p2))
		}
	}
}
