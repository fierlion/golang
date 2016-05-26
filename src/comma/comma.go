package main

import (
    "bytes"
    "fmt"
    "strings"
)

func set_comma(s string) string {
    n := len(s)
    if n <= 3 {
        return s
    }
    return set_comma(s[:n-3]) + "," + s[n-3:]
}

func comma(s string) string {
    n := len(s)
    if n <= 3 {
        return s
    }
    var buf bytes.Buffer
    cur := 0
    if s[cur] == '-' || s[cur] == '+' {
        buf.WriteByte(s[cur])
        cur += 1
    }
    pt := strings.Index(s, ".")
    pre := (pt - cur) % 3
    if pre > 0 {
        fmt.Printf("cur: %d, pre: %d\n", cur, pre)
        buf.WriteString(s[cur:(cur+pre)] + ",")
        cur += pre
    }
    return buf.String()
}

func main() {
    test_a := "123"
    test_b := "1234567.789"
    test_c := "-321123"
    test_d := "-1234567.89123"
    fmt.Printf("%s %s\n", test_a, comma(test_a))
    fmt.Printf("%s %s\n", test_b, comma(test_b))
    fmt.Printf("%s %s\n", test_c, comma(test_c))
    fmt.Printf("%s %s\n", test_d, comma(test_d))
}
