package main

import "fmt"

func checkSign(array []int) string {
    for i := 0; i < len(array)-1; i++ {
        if array[i]*array[i+1] > 0 {
            return "YES"
        }
    }
    return "NO"
}

func main() {
    var n int
    fmt.Print("Введите количество элементов в массиве: ")
    fmt.Scan(&n)

    array := make([]int, n)
    fmt.Print("Введите элементы массива через пробел: ")
    for i := 0; i < n; i++ {
        fmt.Scan(&array[i])
    }

    result := checkSign(array)
    fmt.Println(result)
}