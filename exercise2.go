package main

import (
	"fmt"
)

func main() {

	ii := []int{1,2,3,4,5}}

	n :=foo(ii)
	fmt.Println(n)
}

func foo(xi ...int) int {
	total := o
	for _, v := range xi {
		total += v
	}
	return total
}
