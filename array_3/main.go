package main

import "fmt"

func findMax(array []int) int {
    max := array[0]
    for _, num := range array {
        if num > max {
            max = num
        }
    }
    return max
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

    max := findMax(array)

    fmt.Println("Наибольшее число в массиве:", max)
}
