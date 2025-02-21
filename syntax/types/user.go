package main

import "fmt"

type User struct {
	age  int
	name string
}

func (u User) ChangeName(name string) {
	fmt.Printf("func address %p \n", &u)
	u.name = name
}
func (u *User) ChangeAge(age int) {
	println("func age:", u)
	u.age = age
}
func ChangeUser() {
	u := User{name: "Tom", age: 18}
	fmt.Printf("%+v \n", u)
	fmt.Printf("address %p \n", &u)

	u.ChangeName("Jerry")
	u.ChangeAge(20)
	fmt.Printf("%+v \n", u)
	fmt.Printf("address %p \n", &u)
}
