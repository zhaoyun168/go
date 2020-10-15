package main

import (
	"fmt"
	"log"
)

func defer1()  {
	panic("an error occured: stopping")
	fmt.Println("defer1")
}

func defer2()  {
	defer func() {
		if err := recover();err != nil{
			log.Printf("panic: v%",err)
		}
	}()
	defer1()
	fmt.Println("defer2")
}

func main() {
	fmt.Println("start....")
	defer2()
	fmt.Println("end")

}
