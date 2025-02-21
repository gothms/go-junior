package main

import "fmt"

type Fish struct {
}

func (f Fish) Swim() {

}

// FakeFish 一个全新的类型
type FakeFish Fish

func UseFish() {
	f1 := Fish{}
	f1.Swim()

	f2 := FakeFish{}
	//f2.Swim()

	f3 := FakeFish(f1)
	f4 := Fish(f2)
	fmt.Println(f3, f4)
}

// Yu 别名
type Yu = Fish
