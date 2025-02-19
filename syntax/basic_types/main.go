package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

func main() {
	var a int = 456
	var b int = 123
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(float64(a) / float64(b))
	a++
	b--
	fmt.Println(a, b)

	//var c float64 = 1.23
	//var d int64 = 987
	//fmt.Println(a + c)
	//fmt.Println(a + d)

	fmt.Println(math.Abs(-12.3))

	ExtremeNum()
	String()
	Byte()
	Bool()
}

func ExtremeNum() {
	fmt.Println(math.MaxInt64)
	fmt.Println("float64 最小正数", math.SmallestNonzeroFloat64)
}

func String() {
	// he said `hello go`
	println(`hello, go
换行了！
`)
	println("hello " + strconv.Itoa(123))
	println(len("hello你好"))
	println(utf8.RuneCountInString("hello你好"))
}

func Byte() {
	//var a byte = 12
	var a byte = 'a'
	println(a) // 97

	var str = "hello"
	bs := []byte(str)
	s := string(bs)
	println(s)
}
func Bool() {
	a, b := true, false
	println(a && b)
	println(a || b)
	println(!a)
	//!(a&&b) = !a || !b
	//!(a||b) = !a && !b
}
