package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func swap(x, y *int) {
    tmp := *x
    *x = *y 
    *y = tmp
}

func printArray(arr []int) {
    fmt.Printf("%d", arr[0])
    for i := 1; i<len(arr); i++ {
        fmt.Printf(" %d", arr[i])
    }
    fmt.Println()
}

func partition(arr []int, start, end int) int {
    pivot := arr[end]
    wall := start 
    for i := start; i < end; i++ {
        if (arr[i] <= pivot) { 
            swap(&arr[i], &arr[wall])
            wall += 1
        }
    }
    swap(&arr[wall], &arr[end])
    printArray(arr)
    return wall
}

func quicksort(arr []int, start, end int) {
    if (start >= end) {
        return
    }
    pivot := partition(arr, start, end)
    quicksort(arr, start, pivot-1)
    quicksort(arr, pivot+1, end)
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    numnums_raw, _ := reader.ReadString('\n')
    numnums_trimmed := strings.TrimSpace(numnums_raw)
    numnums, _ := strconv.Atoi(numnums_trimmed)
    nums := make([]int, numnums)
    numarr_raw, _ := reader.ReadString('\n')
    numarr_trimmed := strings.TrimSpace(numarr_raw)
    numarr_split := strings.Split(numarr_trimmed, " ")
    for i, v := range numarr_split {
        num_curr, _ := strconv.Atoi(v)
        nums[i] = num_curr
    }
    quicksort(nums, 0, len(nums)-1)
}

