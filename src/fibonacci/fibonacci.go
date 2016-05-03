/*Given the nth and (n+1)th terms, the (n+2)th can be computed by the 
following relation Tn+2 = (Tn+1)2 + Tn.  Given three integers A, B and N, 
such that the first two terms of the series (1st and 2nd terms) 
are A and B respectively, compute the Nth term of the series. */
package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
    "bufio"
    "math/big"
)

func fib_mod(first, second *big.Int, degree int64) *big.Int {
    two := big.NewInt(int64(2))
    t := new(big.Int)
    squared := t.Exp(second, two, nil)
    u := new(big.Int)
    added := u.Add(first, squared)
    if degree == 0 {
        return u
    }
    return fib_mod(second, added, degree-1)
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    input_raw, _ := reader.ReadString('\n')
    input_trimmed := strings.TrimSpace(input_raw)
    input_split := strings.Split(input_trimmed, " ")
    var inputs[3] int64
    for j, v := range input_split {
        current, _ := strconv.ParseInt(v, 10, 64)
        inputs[j] = current
    }
    fmt.Printf("%d\n", fib_mod(big.NewInt(inputs[0]), big.NewInt(inputs[1]), inputs[2]-3))
}
