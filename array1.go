package main

import "fmt"

func main() {
   var n [10]int /* n 是一个长度为 10 的数组 */
   var i,j int

   /* 为数组 n 初始化元素 */         
   for i = 0; i < 10; i++ {
      n[i] = i + 100 /* 设置元素为 i + 100 */
   }

   /* 输出每个数组元素的值 */
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }

   var m[5]string
   var a,b int

   for a = 0; a < 5; a++ {
      m[a] = fmt.Sprintf("a%d", a) 
   }

   for b = 0; b < len(m); b++ {
      fmt.Println(m[b])   
   }
}