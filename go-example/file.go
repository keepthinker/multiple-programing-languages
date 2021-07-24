package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("types.go", os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		return
	}
	data := make([]byte, 1028)
	count, err := file.Read(data)
	fmt.Printf("count: %d, data: %s", count, data[:count/2])
	fmt.Println()
}
