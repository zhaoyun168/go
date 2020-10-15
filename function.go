package main

import "fmt"

func main() {
	var i, j int = 1, 2
	var max_value int
	max_value = max(i, j)
	fmt.Println(max_value)

	a, b := str("aaa", "bbb")
	fmt.Println(a, b)
}

func max(num1, num2 int) int {
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}

	return result
}

func str(str1, str2 string) (string, string) {
	return str1, str2
}