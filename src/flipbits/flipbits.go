/* flipbits takes a list of 32 bit unsigned ints
   for each it flips the bits in the binary representation
   and returns the the list of resulting unsigned ints
*/

package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
    "bufio"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    raw_tests, _ := reader.ReadString('\n')
    tests, _ := strconv.Atoi(strings.TrimSpace(raw_tests))
    for i := 0; i < tests; i++ {
        num_raw, _ := reader.ReadString('\n')
        num, _ := strconv.ParseUint(strings.TrimSpace(num_raw), 10, 32)
        num_32 := uint32(num)
        fmt.Printf("%b\n", ^num_32)
    }
}
