package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fbnx(i))
	}
}

func fbnx(num int) int {
	if num < 2 {
		return num
	}

	return fbnx(num - 2) + fbnx(num - 1)
}