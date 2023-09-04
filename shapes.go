package main

import "fmt"

type square struct {
	sideLength float64
}

func (square square) getArea() float64 {
	return square.sideLength * square.sideLength
}

type triangle struct {
	height float64
	base   float64
}

func (triangle triangle) getArea() float64 {
	return triangle.height * triangle.base
}

type shape interface {
	getArea() float64
}

func printArea(shape shape) {
	fmt.Println("My area: ", shape.getArea())
}
