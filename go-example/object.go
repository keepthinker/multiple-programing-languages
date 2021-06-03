package main

import "fmt"


var colors = []string {"red", "black", "blue"}

const (
	RED = iota
	BLACK
	BLUE
)

type Animal struct {
	name string 
	weight int
}

type Monkey struct{
	Animal
	color int
}

func (a Monkey) speak() {
	fmt.Printf("%s speak: gigi, weight: %d, color: %s \n", a.name, a.weight, colors[a.color]);
}

func (a Animal) speak() {
	fmt.Printf("%s speak: popo, weight: %d\n", a.name, a.weight)
}


func main() {

	monkey := Monkey{Animal{"monkey", 10}, RED}
	monkey.speak()

	animal := Animal{"bird", 3}
	animal.speak()

}