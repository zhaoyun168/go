package main

import (
	"fmt"
)

func main() {
	fmt.Print("start....")
	panic("an error occured: stopping")
	fmt.Print("end")
}