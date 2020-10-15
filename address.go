package main

import "fmt"

func main() {
   var a int = 10   

   fmt.Printf("变量的地址: %x\n", &a  )

   var b int = 30

   fmt.Printf("变量的地址：%X\n", &b)

   var c *int

   c = &b

   fmt.Println(c)
}