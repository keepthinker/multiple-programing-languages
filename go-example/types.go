package main

import "fmt"
import "errors"

type Color int

func main() {
	var number int8 =  11
	number = 10

	var i, j, k = 0, 1, 2
	fmt.Println(number)
	fmt.Println(i + j + k)

	var a int = 10
	var b int = 3
	var result float32

	result = float32(a) / float32(b)
	fmt.Println(result)

	const c = 20
	fmt.Println(c)

	var d = 3.223423423423
	fmt.Println(d)
	fmt.Println(float32(d))
	fmt.Println(float64(d))

	var str1, str2, str3 = "hello", "world", "!"
	fmt.Println(str1, str2, str3)

	stri1 := "abcdefg"
	strbytes := []byte(stri1)
	strbytes[1] = '2'
	stri1 = string(strbytes)
	fmt.Println("string:", stri1)

	stri2 := "hello"
	fmt.Println("str + str1 = ", stri1 + stri2)

	fmt.Printf("origin: %s, result(2,5): %s\n", stri1, stri1[2:5])

	err := errors.New("file corrupted")
	if err != nil {
		fmt.Println(err)
	}

	const (
		elem1 = iota
		elem2 = iota
		elem3 = iota
		elem4, elem5 = iota, iota
		elem6 = iota
		)
	fmt.Println("elem1=", elem1, "elem2=", elem2, "elem3=", elem3, "elem4=", elem4, "elem5", elem5, "elem6", elem6)

	var h Color = 100
	fmt.Println("color:", h)
}
