package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	"example.com/myutil"
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
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	fmt.Println("start a http server with port 8080")

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/countCharacter", countCharacter)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)

}
