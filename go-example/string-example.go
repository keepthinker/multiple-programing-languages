package main

import "fmt"
import "strings"

func main() {
	var greeting = "hello world!"
	
	fmt.Printf("normal string: ")
	fmt.Printf("normal string: %s\n", greeting)

	const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Printf("quoted string:")
	fmt.Printf("%+q\n", sampleText);

	fmt.Printf("length of greeting: %d\n", len(greeting))
	fmt.Printf(strings.Join([]string{greeting, " join word"}, "|"))
	fmt.Printf("\n")
}
