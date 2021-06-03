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

func (circle *Circle) setRadius(r int) {
	circle.radius = r
}

func sumAndProduct(a int, b int) (int, int) {
	return a + b, a * b
}

func increment(val *int) {
	*val++	
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

	circle3 := Circle{2}
	circle3.setRadius(10)
	fmt.Println("circle3 area:", circle3.area())

	a, b := sumAndProduct(2, 3)
	fmt.Printf("a = %d, b = %d\n", a, b)

	increment(&a);
	fmt.Println("incremented value a:", a)
}
