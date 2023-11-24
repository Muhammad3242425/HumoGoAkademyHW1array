package main

import "fmt"

func reverseArray(array []int) {
    for i := 0; i < len(array)/2; i++ {
        j := len(array) - 1 - i
        array[i], array[j] = array[j], array[i]
    }
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

    reverseArray(array)

    fmt.Println("Переставленный массив:")
    for _, num := range array {
        fmt.Print(num, " ")
    }
}