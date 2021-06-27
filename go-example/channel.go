package main

import (
"fmt"
"time"
)


func main() {
	var ch chan int = make(chan int, 10)
	for i := 0; i < 10; i++ {
		go  transfer(ch, i)
	}
	time.Sleep(time.Second * 2)
	fmt.Println()
	for i := 0; i < 10; i++ {
		val := <-ch
		fmt.Print(val)
	}

}

func transfer(ch chan int, value int) {
	ch <- value
	fmt.Print(value)

}
