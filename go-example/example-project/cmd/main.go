package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/keepthinker/multiple-programing-languages/tree/master/go-example/example-project/internal"
	"github.com/keepthinker/multiple-programing-languages/tree/master/go-example/example-project/internal/traffic"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println(strconv.Atoi("12"))
	traffic.CarStructOperation()

	internal.SliceOperation()
	internal.MapOperation()
	internal.TemplateOperation()

	internal.PanicAndRecover()

	time.Sleep(time.Second * 2)
	fmt.Println("end")

}
