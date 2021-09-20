package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) isAdult() bool {
	return u.age >= 18
}

func (u *user) incrementAge() {
	u.age++
}

func main() {
	u := user{"Alexander", 17}
	u.incrementAge()
	fmt.Print(u, u.isAdult())
}
