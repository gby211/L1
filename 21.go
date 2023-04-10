package main

import "fmt"

type Shape interface {
	draw_shape()
}

// Класс Circle
type Circle struct{}

func (c *Circle) draw_circle() {
	fmt.Println("Рисую круг")
}

// Класс Square
type Square struct{}

func (s *Square) draw_square() {
	fmt.Println("Рисую квадрат")
}

// Адаптер для класса Circle
type CircleAdapter struct {
	Circle *Circle
}

func (ca *CircleAdapter) draw_shape() {
	ca.Circle.draw_circle()
}

// Адаптер для класса Square
type SquareAdapter struct {
	Square *Square
}

func (sa *SquareAdapter) draw_shape() {
	sa.Square.draw_square()
}

// Код, который использует метод draw_shape
func draw_shapes(shapes []Shape) {
	for _, s := range shapes {
		s.draw_shape()
	}
}

func main() {
	circle := &Circle{}
	square := &Square{}

	// Создаем адаптеры для Circle и Square
	circleAdapter := &CircleAdapter{Circle: circle}
	squareAdapter := &SquareAdapter{Square: square}

	// Используем метод draw_shape
	shapes := []Shape{circleAdapter, squareAdapter}
	draw_shapes(shapes)
}
