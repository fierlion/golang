// insertion sort basic implementation
package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "bufio"
)

func print_arr(unsorted []int) {
    fmt.Printf("%d", unsorted[0])
    for j := 1; j < len(unsorted); j++ {
        fmt.Printf(" %d", unsorted[j])
    }
    fmt.Printf("\n")
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    size_raw, _ := reader.ReadString('\n')
    size, _ := strconv.Atoi(strings.TrimSpace(size_raw))
    unsorted := make([]int, size)
    arr_raw, _ := reader.ReadString('\n')
    arr_trimmed := strings.TrimSpace(arr_raw)
    arr_split := strings.Split(arr_trimmed, " ")
    for j, v := range arr_split{
        current, _ := strconv.Atoi(v)
        unsorted[j] = current
    }
    for j := 1; j < len(unsorted); j++ {
        current := unsorted[j]
        for k := 0; k < j; k++ {
            if current < unsorted[k] {
                //make space
                for h := j; h > k; h-- {
                    unsorted[h] = unsorted[h-1]
                }
                unsorted[k] = current
                break
            }
        }
        print_arr(unsorted)
    }
}
