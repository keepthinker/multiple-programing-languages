package main

import (
"fmt"
"sync"
"time"
)

func main() {
	var mutex sync.Mutex

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("try to get lock", i)
			mutex.Lock()
			time.Sleep(time.Millisecond * 100)
			fmt.Println("get the lock", i)
			mutex.Unlock()
		}()
	}
	time.Sleep(time.Second * 2)
}
