package main

import "fmt"
import "math"
import "errors"

func main() {
	var x int = 10
	result, err := sqrt(x)
	if (err == nil) {
		fmt.Printf("sqrt(%d) is %.10f\n", x, result)
	} else {
		fmt.Println("error: ", err)
	}
}

func sqrt(x int) (float64, error){
	if (x < 0) {
		return 0, errors.New("negtive value is not allowed in sqrt")
	}
	return math.Sqrt(float64(x)), nil
}
