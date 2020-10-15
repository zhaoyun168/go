package main

import "fmt"

func main() {
	var a [2]int
	var b [3]int

	a[0] = 111
	b[0] = 123

	fmt.Println("test")

	d := [2]int{11, 22} 
	e := [...]int{11, 22} 
	fmt.Println(d == e)  // 输出 true
	fmt.Println(e[:1])

	f := make([]int, 3, 3)
	f = append(f, 1)
	f = append(f, 2)
	f = append(f, 3)
	fmt.Println(f)

	/*s := make([]int, 5) 
	s = append(s, 6, 7) 
	fmt.Println(len(s), cap(s)) // 输出7 10 
	s = append(s, 8, 9, 10, 11) 
	fmt.Println(len(s), cap(s))//输出11 20 */

	s5 := make([]int, 5) 
	s6 := []int{11, 22, 33, 44, 55, 66} 
	s := append(s5, s6...) 
	fmt.Println(s, len(s), cap(s))//输出[0 0 0 0 0 11 22 33 44 55 66] 11 12 

	data := []int{0,1,2,3,4,5,6,7,8,9}
	var m, n = 3, 9
	mydata := data[m:n]
	r := make([]int, len(mydata))
	copy(r, mydata) 
	fmt.Println(r)

	k := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(k)

	l := make([]int , 3, 3)
	l = append(l, 1)
	l = append(l, 2)
	l = append(l, 3)
	fmt.Println(l)

	mydata1 := l[1:3]

	r1 := make([]int, len(mydata1))

	copy(r1, mydata1)

	fmt.Println(r1)
}