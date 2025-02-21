package main

import (
	"fmt"
	"io"
)

//func Sum[T constraints.Integer](vs []T) T {

func Sum[T Number](vs []T) T {
	var ans T
	for _, v := range vs {
		ans += v
	}
	return ans
}
func Compare[T Number](vs []T, fn func(x T, y T) T) T {
	ans := vs[0]
	for _, v := range vs[1:] {
		ans = fn(ans, v)
	}
	return ans
}
func Find[T any](vs []T, filter func(t T) bool) T {
	for _, v := range vs {
		if filter(v) {
			return v
		}
	}
	var t T
	return t
}
func Insert[T any](vs []T, idx int, val T) []T {
	n := len(vs)
	vs = append(vs, vs[n-1])
	copy(vs[idx+1:n], vs[idx:n-1])
	vs[idx] = val
	return vs
}

type Number interface {
	~int | uint | int32
}
type Int int

func UseSum() {
	sum := Sum[Int]([]Int{1, 2})
	fmt.Println(sum)
}
func Closable[T io.Closer]() {
	var t T
	t.Close()
}

func UseCompare() {
	arr := []int{4, 7, 3, 2, 1, 6, 5}
	minV := Compare(arr, func(x int, y int) int {
		return min(x, y)
	})
	maxV := Compare(arr, func(x int, y int) int {
		return max(x, y)
	})
	fmt.Println(minV, maxV)
}
func UseInsert() {
	arr := []int{1, 2, 3, 5, 6, 7}
	ans := Insert(arr, 3, 4)
	fmt.Println(ans)
}
