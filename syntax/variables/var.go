package main

import (
	"fmt"
	"go-junior/syntax/variables/demo"
)

var (
	a = 123
	b = 12.3
)

func main() {
	fmt.Println(a, b)

	var d uint = 123
	println(d)
	var e int
	println(e)

	println(demo.Global)
	//println(demo.internal)

	println(demo.External)
	//println(demo.external)
}

const (
	State = iota
	State1
	State2

	state6 = 6
)

const (
	A = iota << 1
	B
)
