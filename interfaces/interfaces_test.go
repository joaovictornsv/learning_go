package main

import (
	"math"
	"testing"
)

func TestGeometricFigures(t *testing.T) {
	t.Run("Rect", func(t *testing.T) {
		rect := Rect{10, 20}
		rectArea := rect.area()
		expectedRectArea := float32(10 * 20)

		if rectArea != expectedRectArea {
			t.Fatalf("Rect area error\nExpected: %f\nReceive: %f", expectedRectArea, rectArea)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		circleArea := circle.area()
		expectedCircleArea := float32(math.Pi * 100)

		if circleArea != expectedCircleArea {
			t.Fatalf("Circle area error\nExpected: %f\nReceive: %f", expectedCircleArea, circleArea)
		}
	})

	t.Run("Square", func(t *testing.T) {
		square := Square{20}
		squareArea := square.area()
		expectedSquareArea := float32(math.Pow(20, 2))

		if squareArea != expectedSquareArea {
			t.Fatalf("Square area error\nExpected: %f\nReceive: %f", expectedSquareArea, squareArea)
		}
	})
}
