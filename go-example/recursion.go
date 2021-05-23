package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonaci(i));
	}
	fmt.Println();
}


func fibonaci(i int) int {

	if (i == 0) {
		return 0
	}
	if (i == 1) {
		return 1
	}
	return fibonaci(i - 1) + fibonaci(i - 2)
}


