package main

import (
	"fmt"
)

func sum(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	c <- sum
}

func main() {
	arrs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	c := make(chan int)

	go sum(arrs[:2], c)
	go sum(arrs[2:4], c)
	go sum(arrs[4:6], c)
	go sum(arrs[6:8], c)
	go sum(arrs[8:10], c)

	x1, x2, x3, x4, x5 := <- c, <- c, <- c, <- c, <- c

	fmt.Println(x1+x2+x3+x4+x5)
}