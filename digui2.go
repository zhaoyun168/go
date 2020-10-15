package main

import "fmt"

func main() {
	var result int
	result = jc(10)
	fmt.Println(result)
}

func jc(num int) int {
	if num > 0 {
		return num * jc(num - 1)
	}

	return 1
}