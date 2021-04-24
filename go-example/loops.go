package main

import "fmt"

func main() {
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 6; i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("-------------------");
        var a, b = 0, 10
	for a < b {
		fmt.Println(a)
		a++
	}

	fmt.Println("-------------------");
	for x, y := range arr {
		fmt.Printf("index: %d, value: %d\n", x, y)
	}
}