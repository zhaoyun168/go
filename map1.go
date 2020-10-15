package main

import "fmt"

func main() {
	var arr map[int]int
	arr = make(map[int]int)

	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	arr[4] = 4
	arr[5] = 5
	arr[6] = 6

	for k := range arr {
		fmt.Println("k=>", k, "v=>", arr[k])
	}
}