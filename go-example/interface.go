package main

import "fmt"

type Shape interface {
	area() float64
}
type Rectangle struct {
	width int
	height int
}
type Circle struct {
	radius int
}
func (rectangle Rectangle) area() float64 {
	return float64(rectangle.width * rectangle.height)
}
func (circle Circle) area() float64 {
	return 3.141592654 * float64(circle.radius * circle.radius)
}

func main() {
	var shape Shape = Rectangle{width:4, height:5}
	fmt.Println(shape.area())

	circle := Circle{9}
	shape = circle

	fmt.Println("circle:", shape.area())
	circle.radius = 10

	shape = circle
	fmt.Println("circle:", shape.area())

	var circle2 *Circle
	circle2 = new(Circle)
	(*circle2).radius = 20
	shape = *circle2
	fmt.Println("circle", shape.area())
}
