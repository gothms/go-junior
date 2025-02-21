package main

import "fmt"

func main() {
	arr := make([]int, 256, 767)
	arr = make([]int, 1001, 1884)
	ans, _ := del(arr, 4)
	fmt.Println(len(ans), cap(ans))
	//fmt.Printf("%p,%p \n", arr, ans)
	ans, _ = del(arr, 0)
	fmt.Println(len(ans), cap(ans))
	//fmt.Printf("%p,%p \n", arr, ans)
	ans, _ = del(arr, 2)
	//ans = delete(ans, 2)
	fmt.Println(len(ans), cap(ans))
	//fmt.Printf("%p,%p \n", arr, ans)
}
