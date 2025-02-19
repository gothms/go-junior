package main

import "fmt"

func Defer() {
	defer func() {
		fmt.Println("first defer")
	}()
	defer func() {
		println("second defer")
	}()
}
func DeferClosure() {
	i := 0
	defer func() { // 闭包
		println(i) // 1
	}()
	defer func(i int) {
		println(i) // 0
	}(i)
	i = 1
}

func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a // 0
	// TODO
}
func DeferReturn1() (a int) {
	defer func() {
		a = 1
	}()
	return a // 1
}
func DeferReturn2() *test {
	res := &test{
		name: "Lee",
	}
	defer func() {
		res.name = "Jerry"
	}()
	return res
}

type test struct {
	name string
}

func DeferLoop1() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
func DeferLoop2() {
	for i := 0; i < 10; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}
func DeferLoop3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Println(j)
		}()
	}
}
