package main

import "fmt"

func main() {
   /* 数组长度为 5 */
   var  balance = [5]int {1000, 2, 3, 17, 50}
   var avg float32

   /* 数组作为参数传递给函数 */
   avg = getAverage( balance, 5 ) ;

   /* 输出返回的平均值 */
   fmt.Printf( "平均值为: %f ", avg );

   var arr1 = [6]int{1, 2, 3, 4, 5, 6}
   var avg1 float32

   avg1 = getAverage1(arr1, len(arr1))

   fmt.Printf("avg: %f", avg1)

   a := 1.69
   b := 1.7
   c := a * b      // 结果应该是2.873
   fmt.Println(c)  // 输出的是2.8729999999999998

   a1 := 1690           // 表示1.69
   b1 := 1700           // 表示1.70
   c1 := a1 * b1          // 结果应该是2873000表示 2.873
   fmt.Println(c1)      // 内部编码
   fmt.Println(float64(c1) / 1000000) // 显示
}
func getAverage(arr [5]int, size int) float32 {
   var i,sum int
   var avg float32  

   for i = 0; i < size;i++ {
      sum += arr[i]
   }

   avg = float32(sum) / float32(size)

   return avg;
}

func getAverage1(arr [6]int, size int) float32 {
   var i,sum int
   var avg float32

   for i = 0; i < size; i++ {
      sum += arr[i]
   }

   avg = float32(sum) / float32(size)

   return avg
}