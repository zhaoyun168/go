package main

import "fmt"

var g int

func main() {
	var a, b int = 1, 2

	g = a + b

	fmt.Println(g)
	fmt.Println(add(2, 3))
}

func add(a int, b int) int {
	g = a + b
	return g
}