package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello")
}

func foo() {
	defer func() {
		fmt.Println("der foo")
	}()

	fmt.Println("foo ran")
}
