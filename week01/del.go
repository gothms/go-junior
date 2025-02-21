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
	} else {
		ans = append(arr[:idx], arr[idx+1:]...) // 删除元素
		if ok, c := shrink(cap(arr), n-1); ok { // 是否缩容
			ans = append(make([]T, 0, c), ans...)
		}
	}
	return
}

// shrink 缩容策略
// 参考源码 runtime/slice.go
// func growslice(oldPtr unsafe.Pointer, newLen, oldCap, num int, et *_type) slice
func shrink(c, n int) (ok bool, newC int) {
	const threshold = 256 // 超过 256 后，扩容大小的策略改变
	if c < 32 {           // 容量过小，不考虑
		return
	} else if tt := threshold * 3; c < tt { // 容量 < 768
		if n*3 <= c { // 缩容的空间 >= 扩容的空间
			ok, newC = true, n<<1
		}
	} else if tar := n + (n+tt)/4; tar < c { // 容量 >= 768 且 > 超过 256 后的扩容策略的扩容结果
		if c-tar >= tar-n { // 缩容的空间 >= 扩容的空间
			ok, newC = true, tar
		}
	}
	return
}
