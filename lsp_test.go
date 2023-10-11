package main

import "testing"

/*
Shape Hierarchy with Single Responsibility Principle, Open-Closed Principle, and Liskov Substitution Principle
Checked and reviewed by ChatGPT
*/

func TestPrintShapeInfo(t *testing.T) {
	PrintShapeInfo(&Circle{})
	PrintShapeInfo(&Rectangle{})
}
