package main

import "io"

type Outer struct {
	Inner
}
type Inner struct {
}

func (i Inner) Hello() {
	println("hello~!")
}
func Components() {
	var o Outer
	o.Hello()
}

type OuterPtr struct {
	*Inner
}

func ComponentsPtr() {
	var o Outer
	o.Hello()

	var op OuterPtr
	op.Hello() // panic：指针未初始化
}

type OuterInterface struct {
	io.Closer // 组合接口
}

// SayHello Go没有多态
func (i Inner) SayHello() {
	println("hello,", i.Name())
}
func (o Outer) Name() string {
	return "Outer"
}
func (i Inner) Name() string {
	return "Inner"
}
func UseSayHello() {
	var o Outer
	o.SayHello()
}
