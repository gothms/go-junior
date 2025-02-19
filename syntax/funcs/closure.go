package main

func Closure(name string) func() string {
	// 闭包
	// 返回的这个函数，就是一个闭包
	// 它引用到了 Closure 这个方法的入参
	return func() string {
		return "hello, " + name
	}
}
func ClosureInvoke() {
	c := Closure("Lee")
	println(c())
}
