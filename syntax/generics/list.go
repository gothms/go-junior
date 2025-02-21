package main

type ListGe[T any] interface {
	Add(idx int, val T)
	Append(val T)
	Delete(idx int) error
}
type LinkedListGe[T any] struct {
}

func (l *LinkedListGe[T]) Add(idx int, val T) {

}

func UseList() {
	l := &LinkedListGe[int]{}
	l.Add(1, 123)
}
