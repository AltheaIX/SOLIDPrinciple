package main

import "fmt"

/*
Shape Hierarchy with Single Responsibility Principle, Open-Closed Principle, and Liskov Substitution Principle
Checked and reviewed by ChatGPT
*/

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct{}

func (c *Circle) Area() float64 {
	return 5
}

func (c *Circle) Perimeter() float64 {
	return 10
}

type Rectangle struct{}

func (r *Rectangle) Area() float64 {
	return 3
}

func (r *Rectangle) Perimeter() float64 {
	return 3
}

func PrintShapeInfo(s Shape) {
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}
