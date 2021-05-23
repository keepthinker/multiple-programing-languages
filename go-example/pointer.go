package main

import "fmt"

func sayHello() {
	fmt.Println("hello!");
}
func main() {
	var a int = 0
	var ap *int = &a
	fmt.Printf("address of variable a: %x \n", &a)
	fmt.Printf("variable a: %d \n", *ap)
	*ap = 10
	fmt.Printf("variable a: %d \n", a);

	fmt.Printf("the address of func named hello: %x\n", sayHello);
}
