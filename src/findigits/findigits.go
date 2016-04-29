/* findigits takes as input n numbers.  for each it
   traverses the digits of the number and counts the
   number of times n/digit == 0 
*/
package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    raw_tests, _ := reader.ReadString('\n')
    tests, _ := strconv.Atoi(strings.TrimSpace(raw_tests))
    for i := 0; i < tests; i++ {
        result := 0
        num_raw, _ := reader.ReadString('\n')
        num_str := strings.TrimSpace(num_raw)
        num_full, _ := strconv.Atoi(num_str)
        for j := 0; j < len(num_str); j++ {
            num_cur := int(num_str[j] - '0')
            if num_cur == 0 {
                continue
            }
            modulo := num_full % num_cur
            if modulo == 0 {
                result += 1
            }
        }
        fmt.Printf("%d\n", result)
    }
}
