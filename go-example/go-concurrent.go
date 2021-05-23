package main

import (
	"fmt"
	"time"
)

func say(word string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("say: ", word)
	time.Sleep(100 * time.Millisecond)
}

func sum (arr []int, c chan int) {
	var sum int = 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}
func main() {
	go say("go another routine: hello world")
	say("hello world!")

	numbers := []int {1, 2, 3, 4, 5, 7}
	ch := make(chan int)
	length := len(numbers)
	go sum(numbers[:length / 2], ch)
	go sum(numbers[length / 2:], ch)
	x, y := <-ch, <-ch

	fmt.Printf("x:%d, y:%d, sum:%d\n", x, y, x + y)

}
