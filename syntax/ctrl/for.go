package main

func Loop1() {
	for i := 0; i < 10; i++ {
		println(i)
	}
}
func Loop2() {
	i := 0
	for i < 10 {
		println(i)
		i++
	}
}
func Loop3() {
	for {
		println("hello")
	}
}
func ForArray() {
	println("遍历数组")
	arr := [3]string{"A", "B", "C"}
	for i, s := range arr {
		println(i, s)
	}
}
func ForSlice() {
	println("遍历数组")
	arr := []string{"A", "B", "C"}
	for i, s := range arr {
		println(i, s)
	}
}
func ForMap() {
	m := map[int]string{1: "A", 2: "AB", 3: "ABC"}
	for k, v := range m {
		println(k, v)
	}
}

func LoopBug() {
	users := []user{
		{"TOM"}, {"Jerry"},
	}
	m := make(map[string]*user)
	for _, u := range users {
		m[u.name] = &u
	}
	for k, v := range m {
		println(k, v.name)
	}
}

type user struct {
	name string
}

func ForBreak() {
	i := 0
	for ; ; i++ {
		if i >= 10 {
			break
		}
	}
	println(i)
}
func ForContinue() {
	for i := 0; i < 10; i++ {
		if i&1 == 0 {
			continue
		}
		println(i)
	}
}
