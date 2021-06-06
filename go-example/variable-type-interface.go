package main

import (
	"fmt"
	"strconv"
	"reflect"
)

type Element interface{}

type List []Element;

type Person struct {
	name string
	age int
}

func (p Person) String() string {
	return "name: " + p.name + ", age: " + strconv.Itoa(p.age)
}

func main() {
	list := make(List, 2)
	list[0] = "monkey"
	list[1] = Person{"Kate", 22}
	list = append(list, 12, 3.23, "John")

	for i, elem := range list {
		if value, ok := elem.(string); ok {
			fmt.Printf("index: %d, value: %s\n", i, value)
		}
		if value, ok := elem.(int); ok {
			fmt.Printf("index: %d, value: %d\n", i, value)
		}
		if value, ok := elem.(Person); ok {
			fmt.Printf("index: %d, value: %s\n", i, value)
		}
		if value, ok := elem.(float64); ok {
			fmt.Printf("index: %d, value: %f\n", i, value)
		}
	}

	fmt.Println("\nswitch exmple:")

	for i, elem := range list {
		switch value := elem.(type) {
		case string :
			fmt.Printf("index: %d, value: %s\n", i, value)
		case int :
			fmt.Printf("index: %d, value: %d\n", i, value)
		case Person :
			fmt.Printf("index: %d, value: %s\n", i, value)
		case float64:
			fmt.Printf("index: %d, value: %f\n", i, value)
		}
	}

	var floatNum float64 = 6.4123

	t := reflect.TypeOf(floatNum)
	v := reflect.ValueOf(floatNum)

	fmt.Println("type:", t)
	fmt.Println("value1:", v)

	p := reflect.ValueOf(&floatNum)
	v2 := p.Elem()
	v2.SetFloat(7.3)
	fmt.Println("value2:", v2)
	fmt.Println("value1:", v)

}
