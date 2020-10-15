package main

 

import (

"fmt"

)

 

func main() {

arr := []int{1, 2, 3}

fmt.Println("cap(arr) : ", cap(arr))

fmt.Println("len(arr) : ", len(arr))

 

//长度和容量都为5

slice1 := make([]string, 5)

//长度为3，容量为5

slice2 := make([]int, 3, 5)

fmt.Println("cap(slice1) : ", cap(slice1))

fmt.Println("cap(slice2) : ", cap(slice2))

 

c1 := make(chan int)

c2 := make(chan int, 2)

fmt.Println("cap(c1) : ", cap(c1))

fmt.Println("cap(c2) : ", cap(c2))

}