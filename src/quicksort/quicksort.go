package main

import (
    "fmt"
)

func swap(x, y *int) {
    tmp := *x
    *x = *y 
    *y = tmp
}

func partition(arr []int, start, end int) int {
    pivot := arr[start]
    wall := start
    for i := start+1; i<= end; i++ {
        if (arr[i] < pivot) { 
            wall += 1
            swap(&arr[i], &arr[wall])
        }
    }
    swap(&arr[start], &arr[wall])
    return wall
}

func quicksort(arr []int, start, end int) {
    if (start >= end) {
        return
    }
    pivot := partition(arr, start, end)
    quicksort(arr, start, pivot)
    quicksort(arr, pivot+1, end)
}

func main() {
    s := []int{6,5,4,3,2,1}
    quicksort(s, 0, len(s)-1)
    fmt.Printf("%v\n", s)
}

