package main

import (
"fmt"
)

type getNumberType func(int) bool

func isPositive(num int) bool {
	return num >= 0
}

func isNegtive(num int) bool {
	return num < 0
}

func printNumType(val int, judgeNumType getNumberType) {
	var a string
	if judgeNumType(val) {
		a = "true"
	} else {
		a = "false"
	}
	fmt.Printf("val:%d is %s\n", val, a)
}

func main() {
	fmt.Printf("positive: ")
	printNumType(10, isPositive)
	fmt.Printf("negtive: ")
	printNumType(11, isNegtive)
	fmt.Printf("positive: ")
	printNumType(-88, isPositive)

}
