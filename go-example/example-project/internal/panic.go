package internal

import "fmt"

func PanicAndRecover() {
	defer func() {
		p := recover()
		fmt.Printf("recover from panic, %s\n", p)
	}()

	panic("error panic")
}

func PanicExit() {
	panic("panic, programe crashes")
}

func PanicRecoverPanicExit() {
	defer func() {
		p := recover()
		panic("can't recover | " + fmt.Sprintf("%s", p))
	}()
	panic("panic, programe crashes")
}
