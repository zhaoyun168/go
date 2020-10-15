package main

import "fmt"
import "unsafe"

func main() {
	const STR string = "abc"
	fmt.Println(STR)

	const (
		Unknown = 0
	    Female = 1
	    Male = 2
	)

	fmt.Println(Male)

	fmt.Println(unsafe.Sizeof(STR))

	var a, b int = 10, 20
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	if ( a > b ) {
		fmt.Println("a > b")
	} else {
		fmt.Println("a <= b")
	}

	var c, d bool = true, false
	if ( c && d || c) {
		fmt.Println("c")
	} else {
		fmt.Println("d")
	}

	var k string = "123"
	j := &k
	fmt.Println(j)
	fmt.Println(*j)
}