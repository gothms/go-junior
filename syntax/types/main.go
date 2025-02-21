package main

import "fmt"

func main() {
	u1 := &User{}
	println(u1)
	u1 = new(User)
	println(u1)

	u2 := User{}
	fmt.Printf("%+v\n", u2)
	u2.name = "Jerry"
	println(u2.name)

	var u3 User
	fmt.Printf("%+v\n", u3)
	var u4 *User
	println(u4)

	ChangeUser()

	//ComponentsPtr()

	UseSayHello()
}
