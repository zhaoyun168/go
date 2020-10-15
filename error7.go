package main

import (
	"fmt"
	"log"
)

func divnum(a int, b int) int {
	//捕获异常信息
	defer func() {
		if err := recover();err != nil {
			log.Printf("异常信息[%v]", err)
		}
	}()

	divv := a / b

	return divv
}

func main() {
	var a, b int = 30, 0
	value := divnum(a, b)
	fmt.Println(value)
}