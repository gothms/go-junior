package main

import "fmt"

// delete 删除操作
func delete[T any](arr []T, idx int) []T {
	if idx < 0 || idx >= len(arr) {
		return arr
	}
	// 写法一：拷贝
	//copy(arr[idx:], arr[idx+1:])
	//return arr[:len(arr)-1]
	// 写法二：append
	return append(arr[:idx], arr[idx+1:]...)
}

// del 删除操作，且支持缩容
// 参考源码 slices/slices.go
// func Delete[S ~[]E, E any](s S, i, j int) S
func del[T any](arr []T, idx int) (ans []T, err error) {
	n := len(arr)
	if idx < 0 || idx >= n { // 索引不合法
		ans, err = nil, fmt.Errorf("illegal parameter, index out of range [%d] with length %d", idx, n)
		return
	}
	if ok, c := shrink(cap(arr), n-1); ok { // 是否缩容
		ans = make([]T, idx, c) // 缩容
		copy(ans, arr[:idx])
		ans = append(ans, arr[idx+1:]...) // 删除元素
	} else {
		ans = append(arr[:idx], arr[idx+1:]...) // 删除元素
	}
	return
}

// shrink 缩容策略
// 示例一：len=256，cap=767，缩容后cap=510。缩减了 256 个元素空间，同时还可以在不扩容的情况下添加 255 个元素
// 示例二：len=1001，cap=1884，缩容后cap=1442。缩减了 442 个元素空间，同时还可以在不扩容的情况下添加 442 个元素
// 参考源码 runtime/slice.go
// func growslice(oldPtr unsafe.Pointer, newLen, oldCap, num int, et *_type) slice
func shrink(c, n int) (ok bool, newC int) {
	if c < 32 { // 容量过小，不考虑
		return
	}
	const threshold = 256            // 超过 256 后，扩容大小的策略改变
	if tt := threshold * 3; c < tt { // 1.容量 < 768
		if n*3 <= c { // 缩容的空间 >= 扩容的空间
			ok, newC = true, n<<1
		}
	} else if tar := n + (n+tt)/4; tar < c { // 2.容量 >= 768 且 > 超过 256 后的扩容策略的扩容结果
		if c-tar >= tar-n { // 缩容的空间 >= 扩容的空间
			ok, newC = true, tar
		}
	}
	return
}
