package internal

import (
	"fmt"

	"github.com/keepthinker/multiple-programing-languages/tree/master/go-example/example-project/internal/traffic"
)

func MapOperation() {
	m := make(map[int]*traffic.Car)

	m[0] = nil

	v, ok := m[0]
	fmt.Printf("map key:%d ok:%v value:%v\n", 0, ok, v)

	m = map[int]*traffic.Car{
		0: &(traffic.Car{Name: "Porsche"}),
		1: &(traffic.Car{Name: "BWM"}),
	}
	v, ok = m[0]
	fmt.Printf("map key:%d ok:%v value:%v\n", 0, ok, v)

	delete(m, 0)
	v, ok = m[0]
	fmt.Printf("map key:%d ok:%v value:%v\n", 0, ok, v)
}
