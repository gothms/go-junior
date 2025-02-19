package main

func main() {
	Invoke()
}
func Func0(name string) string {
	return "hello, " + name
}
func Func1(a, b, c int, d string) (string, error) {
	return "hello world, ", nil
}
func Func2(a, b int) (str string, err error) {
	str = "hello world, "
	return
}
func Func3(a, b int) (str string, err error) {
	return "abc", nil
}

func Recursive() {
	//Recursive()
}

func Invoke() {
	s := Func0("Lee")
	println(s)

	str, err := Func2(1, 2)
	println(str, err)

	_, err = Func3(3, 4)

	Recursive()

	Func4()
	Func5()
	Func6Invoke()
	Func7()

	ClosureInvoke()

	YourNameInvoke()

	Defer()
	DeferClosure()
	a := DeferReturn()
	println("a:", a)
	a = DeferReturn1()
	println("a:", a)
	res := DeferReturn2()
	println("name:", res.name)

	DeferLoop1()
	DeferLoop2()
	DeferLoop3()
}
