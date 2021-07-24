package main

import "fmt"

type say func() string

func sayHello() string {
	return "hello"
}

func main() {

	fmt.Printf("b = %t\n", throwPanic(sayHello))
	fmt.Printf("b = %t\n", throwsPanic2(func(){ panic("always panic") }))
}

func throwPanic(f say) (b bool) {
	b = true
	f()
	return b
}

func throwsPanic2(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}
