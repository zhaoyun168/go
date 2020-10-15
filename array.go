package main

import "fmt"

func main() {
	var array = []int{1,2,3,4,5,6}

	array[1] = 22

	for k, v := range array {
		fmt.Println(k, v)
	}

	var arr = []string{"a", "b", "c", "d", "e", "f", "g"}

	for k, v := range arr {
		fmt.Println(k, v)
	}
}