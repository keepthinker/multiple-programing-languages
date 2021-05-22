package main

import "fmt"

func main() {
	fmt.Println("hello world!");
	numbers := []int{1, 2, 3, 4, 5, 6, 7 ,8 ,9}

	for i := range numbers {
		fmt.Printf("number: %d\n", numbers[i])
	}

	nameSalaryMap := map[string] int {"Shawn": 100000, "John": 80000 }
	
	for key := range nameSalaryMap {
		fmt.Printf("name:%s, salary:%d\n", key, nameSalaryMap[key])
	}

	for key, value := range nameSalaryMap {
		fmt.Println("name:", key, " salary:", value)
	}

}
