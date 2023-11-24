package main

import (
 "fmt"
)

func main() {
 var n int
 fmt.Scan(&n)

 numbers := make([]int, n)
 for i := 0; i < n; i++ {
  fmt.Scan(&numbers[i])
 }

 minOdd := int(1e9) // Инициализируем переменную с очень большим значением

 for _, num := range numbers {
  if num%2 != 0 && num < minOdd {
   minOdd = num
  }
 }

 if minOdd == int(1e9) {
  minOdd = 0
 }

 fmt.Println(minOdd)
}
