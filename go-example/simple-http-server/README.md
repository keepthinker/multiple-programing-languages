# Create module
## example
```shell
mkdir greetings
cd greetings
go mod init example.com/greetings
touch greetings.go
```

edit source code of greetings.go

```go
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

```shell
cd ..

mkdir hello

cd hello

go mod init example.com/hello
```

edit source code of hello.go

```go
package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
```
go mod edit -replace=example.com/greetings=../greetings
go mod tidy
go run .

# References
[golang-create-module](https://golang.org/doc/tutorial/create-module)
