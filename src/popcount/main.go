package main

import (
    "fmt"
    "popcount1"
    "popcount2"
    "popcount3"
    "popcount4"
)

func main() {
    test_int := uint64(5)
    test_int_lg := uint64(1<<64-1)
    test_res1 := popcount1.PopCount(test_int)
    test_res1_lg := popcount1.PopCount(test_int_lg)
    test_res2 := popcount2.PopCount(test_int)
    test_res2_lg := popcount2.PopCount(test_int_lg)
    test_res3 := popcount3.PopCount(test_int)
    test_res3_lg := popcount3.PopCount(test_int_lg)
    test_res4 := popcount4.PopCount(test_int)
    test_res4_lg := popcount4.PopCount(test_int_lg)
    fmt.Printf("%d\n", test_res1)
    fmt.Printf("%d\n", test_res1_lg)
    fmt.Printf("%d\n", test_res2)
    fmt.Printf("%d\n", test_res2_lg)
    fmt.Printf("%d\n", test_res3)
    fmt.Printf("%d\n", test_res3_lg)
    fmt.Printf("%d\n", test_res4)
    fmt.Printf("%d\n", test_res4_lg)
}
