package main

import "fmt"

type Square struct {
	side_length float64
}

func main() {
	var s1 Square
	s1.side_length = 100.00
	fmt.Println(s1.getArea())
}

func (s Square) getArea() float64 {
	return s.side_length * s.side_length
}