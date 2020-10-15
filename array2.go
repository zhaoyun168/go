package main

import "fmt"

func main() {
	var arr = [2][2]int{
		{1, 2},
		{3, 4},
	}

	fmt.Println(arr)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arr[i][j])	
		}
	}

	var array = [3][3]string{
		{"a1", "b1", "c1"},
		{"a2", "b3", "c3"},
		{"a3", "b3", "c3"},
	}

	fmt.Println(array)

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Println(array[i][j])	
		}
	}
}