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

func fibonacci(n int, ch chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func fibonacci2(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
			case c <- x:
				x, y = y, x + y
			case <- quit:
				fmt.Println("end!")
				return

		}
	}
}
func main() {
	go say("go another routine: hello world")
	say("hello world!")

	numbers := []int {1, 2, 3, 4, 5, 6}
	ch := make(chan int)
	length := len(numbers)
	go sum(numbers[:length / 2], ch)
	go sum(numbers[length / 2:], ch)
	x, y := <-ch, <-ch

	fmt.Printf("x:%d, y:%d, sum:%d\n", x, y, x + y)

	ch = make(chan int, 10)
	go fibonacci(10, ch)
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("select example")
	var numCh chan int = make(chan int)
	quitCh := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-numCh)
		}
		quitCh <- 0
	}()
	fibonacci2(numCh, quitCh)

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for {
			select {
				case x := <-c1:
				fmt.Printf("c1 value:%d\n", x)
				case <-time.After(time.Second * 3):
				fmt.Println("timeout")
				c2 <- 0
			}
		}
	}()

	go func() {
		<-time.After(time.Second * 2)
		c1 <- 100
	}()
	c1 <- 1000000000
	<-c2

	fmt.Println("finished")
}
