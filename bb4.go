package main

import (
	"fmt"
)

func bb() func(f int) int {
	var a int = 10
	return func(f int) int {
		a += f
		return a
	}
}

func main() {
	c := bb()
	fmt.Println(c(2))
}