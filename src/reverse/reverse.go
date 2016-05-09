package main

import (
    "fmt"
)

func reverse(arr []int) {
    for i := 0; i < len(arr)/2; i++ {
        end := (len(arr) - i) - 1
        arr[i], arr[end] = arr[end], arr[i]
    }
}

func main() {
    s := []int{1,2,3,4,5}
    s1 := []int{0}
    s2 := []int{1,2,3,4,5,6}
    reverse(s)
    reverse(s1)
    reverse(s2)
    fmt.Printf("%v\n", s)
    fmt.Printf("%v\n", s1)
    fmt.Printf("%v\n", s2)
}
