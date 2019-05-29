package main

import (
	"fmt"
)

func main() {
	//Exercis e 4.3: Rewrite reverse to use an array pointer instead of a slice.
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(array) //[1 2 3 4 5 6 7 8]
	Reverse(&array)
	fmt.Println(array) // [8 7 6 5 4 3 2 1]
	//Exercis e 4.4: Write a version of rotate that operates in a single pass.
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)
	RotateToTheLeft(slice, 6)
	fmt.Println(slice)

	str := []byte("\x48\x65\x6C\x6C\x6F\x2C\x20\xe4\xb8\x96\xe7\x95\x8c")
	fmt.Println(string(str)) //Hello, 世界
	ReverseUTF(str)
	fmt.Println(string(str)) //界世 ,olleH
}
