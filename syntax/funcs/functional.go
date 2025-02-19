package main

import "fmt"

func Func4() {
	f := Func3
	s, err := f(1, 2)
	println(s, err)
}
func Func5() {
	fn := func(a, b int) int {
		return a + b
	}
	i := fn(1, 21)
	fmt.Println(i)
}
func Func6() func(name string) string {
	return func(name string) string {
		return "hello, " + name
	}
}
func Func6Invoke() {
	fn := Func6()
	str := fn("Lee")
	println(str)
}
func Func7() {
	fn := func(name string) string {
		return "hello, " + name
	}("Lee")
	println(fn)
}
