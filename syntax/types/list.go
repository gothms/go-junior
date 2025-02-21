package main

type List interface {
	Add(idx int, val any)
	Append(val any)
	Delete(idx int) error
}
type LinkedList struct {
	head node
}

func (l *LinkedList) Add(idx int, val any) {

}

type node struct {
}

func UseList() {
	l := &LinkedList{}
	l.Add(1, 123)
	l.Add(1, "123")
	l.Add(1, nil)
}
