// stringrid finds the ocurrence of a given 2D pattern
// of digits in a larger 2D array.
package main

import (
    "bufio"
    "os"
    "fmt"
    "strconv"
    "strings"
)

func find_grid(large []string, small []string, cols int, offset int, found bool) bool {
    for j := range large {
        start_index := strings.Index(large[j][offset:], small[0])
        if (start_index != -1) {
            //found one..try other small_arr cols
            found = true
            for k := 1; k < cols; k++ {
                if (j+k < len(large)) {
                    cur_index := strings.Index(large[j + k][offset:], small[k])
                    if (cur_index != start_index) {
                        found = false
                        break
                    }
                }
            }
            if (found == true) {
                return true
            }
        }
    }
    return false
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    raw_tests, _ := reader.ReadString('\n')
    num_tests, _ := strconv.Atoi(strings.TrimSpace(raw_tests))
    for i := 0; i < num_tests; i++ {
        found_it := false
        arr_size_raw, _ := reader.ReadString('\n')
        trimmed := strings.TrimSpace(arr_size_raw)
        split := strings.Split(trimmed, " ")
        cols, _ := strconv.Atoi(split[0])
        rows, _ := strconv.Atoi(split[1])
        large_arr := make([]string, cols)
        for j := 0; j < cols; j++ {
            row_raw, _ := reader.ReadString('\n')
            trimmed_row := strings.TrimSpace(row_raw)
            large_arr[j] = trimmed_row
        }
        sarr_size_raw, _ := reader.ReadString('\n')
        strimmed := strings.TrimSpace(sarr_size_raw)
        ssplit := strings.Split(strimmed, " ")
        scols, _ := strconv.Atoi(ssplit[0])
        srows, _ := strconv.Atoi(ssplit[1])
        small_arr := make([]string, scols)
        for j := 0; j < scols; j++ {
            row_raw, _ := reader.ReadString('\n')
            trimmed_row := strings.TrimSpace(row_raw)
            small_arr[j] = trimmed_row
        }
        for ofst := 0; ofst < ((rows - srows)+1); ofst ++ {
            found_it = find_grid(large_arr, small_arr, scols, ofst, found_it)
            if found_it == true {
                break
            }
        }
        if (found_it == true) {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}
