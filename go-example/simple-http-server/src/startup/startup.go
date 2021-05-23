package main

import (
	"example.com/myutil"
	"fmt"
	"net/http"
	"strconv"
)

func hello(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprint(writer, "hello world")
}

func countCharacter(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	word := values.Get("word")
	character := values.Get("character")
	sum := myutil.CountCharacter(word, int32(character[0]))

	writer.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(writer, strconv.Itoa(sum))
}

func main() {
	fmt.Println("start a http server with port 8080")

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/countCharacter", countCharacter)

	http.ListenAndServe(":8080", nil)

}
