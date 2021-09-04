package internal

import (
	"fmt"
)

func SliceOperation() {
	fmt.Println("compare: ")
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	// d := [3]int{1, 2}
	// fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
	subSlice()

	arr := [...]int{0, 1, 2, 3, 4, 5}
	arr1 := arr[:]
	reverse(arr1)
	fmt.Println(arr)

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func subSlice() {
	slice := []int{1, 2, 3, 4, 5}
	printSlice(slice)
	newSlice := slice[1:3]
	newSlice[1] = 100
	printSlice(newSlice)
	printSlice(slice)
}

func printSlice(slice []int) {
	fmt.Printf("-----print slicc-----len:%d----\n", len(slice))
	for i, v := range slice {
		fmt.Printf("index:%d, value:%d\n", i, v)
	}
}
