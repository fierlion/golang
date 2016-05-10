package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func printArray(ar []int) {
    if (len(ar) > 0) {
        fmt.Printf("%d", ar[0])
        for i := 1; i<len(ar); i++ {
            fmt.Printf(" %d", ar[i])
        }
        fmt.Println()
    }
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    raw_length, _ := reader.ReadString('\n')
    trimmed_length := strings.TrimSpace(raw_length)
    numlength, _ := strconv.Atoi(trimmed_length)
    raw_nums, _ := reader.ReadString('\n')
    trimmed_nums := strings.TrimSpace(raw_nums)
    split_nums := strings.Split(trimmed_nums, " ")
    var nums = make([]int, numlength)
    for i, v := range split_nums {
        curr, _ := strconv.Atoi(v)
        nums[i] = curr 
    }
    raw_length, _ = reader.ReadString('\n')
    trimmed_length = strings.TrimSpace(raw_length)
    numlength, _ = strconv.Atoi(trimmed_length)
    raw_nums, _ = reader.ReadString('\n')
    trimmed_nums = strings.TrimSpace(raw_nums)
    split_nums = strings.Split(trimmed_nums, " ")
    var concord [100010]int
    for _, v := range split_nums {
        curr, _ := strconv.Atoi(v)
        concord[curr] += 1
    }
    for _, v := range nums {
        concord[v] -= 1
    }
    result := make([]int, 0)
    for k, v := range concord {
        if (v > 0) {
            result = append(result, k)
        } 
    }
    printArray(result)
}
