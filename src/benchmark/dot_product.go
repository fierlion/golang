package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "encoding/binary"
    "sync"
)

func dotprod(input1, input2, output [5][5]uint32, start, end int, wg *sync.WaitGroup) {
    for i := start; i < end; i++ {
        for j := start; j < end; j++ {
            fmt.Printf("%d, %d, %d\n", i, j, input1[i][j])
            fmt.Printf("%d, %d, %d\n", i, j, input2[i][j])
        }
    }
    wg.Done()
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    file1_raw, _ := reader.ReadString('\n')
    file1_stripped := strings.TrimSpace(file1_raw)
    f1, _ := os.Open(file1_stripped)
    file2_raw, _ := reader.ReadString('\n')
    file2_stripped := strings.TrimSpace(file2_raw)
    f2, _ := os.Open(file2_stripped)
    var ints1 [5][5]uint32
    var ints2 [5][5]uint32
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            b1 := make([]byte, 4)
            b2 := make([]byte, 4)
            f1.Read(b1)
            f2.Read(b2)
            this_int1 := binary.LittleEndian.Uint32(b1)
            this_int2 := binary.LittleEndian.Uint32(b2)
            ints1[i][j] = this_int1
            ints2[i][j] = this_int2
        }
    }
    var results [5][5]uint32
    threads_raw, _ := reader.ReadString('\n')
    threads_stripped := strings.TrimSpace(threads_raw)
    threads, _ := strconv.Atoi(threads_stripped)
    var wg sync.WaitGroup
    wg.Add(threads)
    for t := 0; t < threads; t ++ {
        go_start := t * (5/threads)
        go_end := (t+1) * (5/threads)
        go dotprod(ints1, ints2, results, go_start, go_end, &wg)
    }
    wg.Wait()
}
