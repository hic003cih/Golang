package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (u person) speak() {
	fmt.Println(u.first, u.last, u.age)
}
func main() {
	leif := person{
		first: "leif",
		last:  "chang",
		age:   15,
	}

	leif.speak()
}
