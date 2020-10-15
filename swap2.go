package main

import "fmt"

func main() {
	//var a, b int = 10, 20
	a, b := 10, 20

	fmt.Printf("before a: %d\n", a)
	fmt.Printf("before b: %d\n", b)

	swap(&a, &b)

	fmt.Printf("after a: %d\n", a)
	fmt.Printf("after b: %d\n", b)
}

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}