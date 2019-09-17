package main

import "fmt"

/* func Repeat(s string) string {

	var s2 string

	for i := 0; i < 5; i++ {
		s2 += "aa"
	}
	return s2
}
 */

 //重構,將重複次數用一個常數拉出來,以後就不用到for裡面更改次數

const repeatCount = 5

// Repeat returns character repeated 5 times
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func main()  {
	fmt.Println(Repeat("X"))
}