package main

import "fmt"
//import "time"

func main() {
	/*fmt.Println("main start")

	go func() {
		fmt.Println("goroutine")
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("main end")*/

	ch1 := make(chan string, 10)

	ch1 <- "a"
	ch1 <- "b"

	val, ok := <- ch1
	val1, ok1 := <- ch1

	fmt.Println(val, ok)
	fmt.Println(val1, ok1)
}