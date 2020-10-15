package main

import (
	"fmt"
)

func defer1()  {
	panic("an error occured: stopping")
	fmt.Println("defer1")
}

func defer2()  {
	defer1()
	fmt.Println("defer2")
}

func main() {
	fmt.Println("start....")
	defer defer2()
	fmt.Println("end")

}