package main

import (
  "crypto/sha256"
  "fmt"
)

func sha_diff(first, second [32]byte) int {
  result := 0
  for k := 0; k < 32; k++ {
    for i := 0; i < 6; i++ {
      firstBit := first[k] & (1 << uint(i)) > 0
      secBit := second[k] & (1 << uint(i)) > 0
      if firstBit != secBit {
        result += 1
      }
    }
  }
  return result
}

func main() {
  c1 := sha256.Sum256([]byte("x"))
  c2 := sha256.Sum256([]byte("x"))
  result := sha_diff(c1, c2)
  fmt.Printf("%x\n%x\n%d\n", c1, c2, result)
}
