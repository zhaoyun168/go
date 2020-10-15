package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i <= 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	strings := []string{"a1", "a2", "a3"}

	for i, v := range strings {
		fmt.Println(i, v)
	}

	ints := [6]int{1, 2, 3, 4, 5}

	for i, v := range ints {
		fmt.Println(i, v)
	}
}