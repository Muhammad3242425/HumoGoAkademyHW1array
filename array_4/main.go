package main

import "fmt"

func countElements(array []int) int {
    count := 0
    n := len(array)

    for i := 1; i < n-1; i++ {
        if array[i] > array[i-1] && array[i] > array[i+1] {
            count++
        }
    }

    return count
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

    count := countElements(array)

    fmt.Println("Количество элементов, у которых два соседних элемента меньше данного:", count)
}
