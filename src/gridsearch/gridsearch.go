// gridsearch finds the ocurrence of a given 2D pattern
// of digits in a larger 2D array.
// NOTE: I am just learning go so please forgive me this code...
package main

import (
    "bufio"
    "os"
    "fmt"
    "strconv"
    "strings"
)

func check_grids(small []int, large []int, col int, row int, scol int, srow int) bool{
    for i := range small {
    }
}


func main() {
    reader := bufio.NewReader(os.Stdin)
    raw_tests, _ := reader.ReadString('\n')
    num_tests, _ := strconv.Atoi(strings.TrimSpace(raw_tests))
    found_it := false
    for i := 0; i < num_tests; i++ {
        arr_size_raw, _ := reader.ReadString('\n')
        trimmed := strings.TrimSpace(arr_size_raw)
        split := strings.Split(trimmed, " ")
        cols, _ := strconv.Atoi(split[0])
        rows, _ := strconv.Atoi(split[1])
        full_arr := make([]int, cols*rows)
        for j := 0; j < cols; j++ {
            row_raw, _ := reader.ReadString('\n')
            trimmed_row := strings.TrimSpace(row_raw)
            split_row := strings.Split(trimmed_row, "")
            for k, v := range split_row {
                curr_num, _ := strconv.Atoi(v)
                full_arr[j*cols + k] = curr_num
            }
        }
        sarr_size_raw, _ := reader.ReadString('\n')
        strimmed := strings.TrimSpace(sarr_size_raw)
        ssplit := strings.Split(strimmed, " ")
        scols, _ := strconv.Atoi(ssplit[0])
        srows, _ := strconv.Atoi(ssplit[1])
        small_arr := make([]int, scols*srows)
        for j := 0; j < scols; j++ {
            row_raw, _ := reader.ReadString('\n')
            trimmed_row := strings.TrimSpace(row_raw)
            split_row := strings.Split(trimmed_row, "")
            for k, v := range split_row {
                curr_num, _ := strconv.Atoi(v)
                small_arr[j*scols + k] = curr_num
            }
        }
        for j := 0; j < (cols-scols)+1; j++ {
            for k := 0; k < (rows-srows) +1; k++ {
                if small_arr[0] == full_arr[j*cols + k] {
                   check_grids(small_arr, full_arr, j, k, scols, srows)
                   //fmt.Printf("small: %d, large %d\n", small_arr[0], full_arr[j*cols + k])
                }
            }
        }
    }
}
