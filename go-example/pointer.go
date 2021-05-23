package main

import "fmt"

func main() {
	var a int = 0
	var ap *int = &a
	fmt.Printf("address of variable a: %x \n", &a)
	fmt.Printf("variable a: %d \n", *ap)
	*ap = 10
	fmt.Printf("variable a: %d \n", a);
}
