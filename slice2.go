package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5,6}

	fmt.Println("1-3:", arr[:3])
	fmt.Println("2-5:", arr[1:5])
	fmt.Println("2-:", arr[1:])

	var arr1 []int

	arr1 = append(arr1, 0)
	arr1 = append(arr1, 1)
	arr1 = append(arr1, 2)
	arr1 = append(arr1, 3)
	arr1 = append(arr1, 4)
	arr1 = append(arr1, 5)
	arr1 = append(arr1, 6)

	fmt.Println(arr1)
}