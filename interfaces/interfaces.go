package main

import (
	"fmt"
	"math"
)

type GeometricFigure interface {
	area() float32
}

type Rect struct {
	heigth float32
	width  float32
}

func (r Rect) area() float32 {
	return r.heigth * r.width
}

type Circle struct {
	radius float32
}

func (c Circle) area() float32 {
	return float32(math.Pi * math.Pow(float64(c.radius), 2))
}

type Square struct {
	side float32
}

func (s Square) area() float32 {
	return float32(math.Pow(float64(s.side), 2))
}

func escreverArea(figure GeometricFigure) {
	fmt.Printf("A área da figura é: %0.2f\n", figure.area())
}

func main() {
	rect := Rect{10, 20}
	circle := Circle{15}
	square := Square{8}

	escreverArea(rect)
	escreverArea(circle)
	escreverArea(square)
}
