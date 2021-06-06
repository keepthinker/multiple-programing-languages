package main

import ("fmt"
	"strconv"
)

const Pi float64 = 3.141592654

type Shape interface {
	area() float64
	perimeter() float64
	addArea(area int) float64
}
type Rectangle struct {
	width int
	height int
}
type Circle struct {
	radius int
}
func (r Rectangle) area() float64 {
	return float64(r.width * r.height)
}

func (r Rectangle) perimeter() float64 {
	return float64(r.width * r.height)
}

func (r Rectangle) addArea(area int) float64 {
	return float64(area) + r.area()
}

func (r Rectangle) String() string {
	return "regtangle, width: " + strconv.Itoa(r.width) + ", height: " + strconv.Itoa(r.height)
}

func (c Circle) area() float64 {
	return Pi * float64(c.radius * c.radius)
}

func (c Circle) perimeter() float64 {
	return float64(2 * c.radius) * Pi
}

func (c Circle) addArea(area int) float64 {
	return float64(area) + c.area()
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

	fmt.Printf("circle3 add area %d, circle3.area: %f\n", 10, circle3.addArea(10))

	a, b := sumAndProduct(2, 3)
	fmt.Printf("a = %d, b = %d\n", a, b)

	increment(&a);
	fmt.Println("incremented value a:", a)
	rec := Rectangle{10, 20}
	fmt.Println("rec:", rec)
}