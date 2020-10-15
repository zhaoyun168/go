package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func main() {
	//记录开始时间
	start := time.Now().UnixNano()
	begin := time.Now()

	var s [1000000]int

	total := 0
	for i := 0; i < 1000000; i++ {
		s[i] = i	
	}

	for _, v := range s {
		total += v
	}

	fmt.Println(total)

	//记录结束时间
	end := time.Now().UnixNano()

 
	//输出执行时间，单位为毫秒。
	fmt.Println((end - start))
	
	// work
	fmt.Println("time: ", time.Since(begin))

}