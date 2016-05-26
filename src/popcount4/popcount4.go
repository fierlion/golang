package popcount4

func PopCount(x uint64) int {
    result := 0
    for x != 0 {
        result += 1
        x = x&(x-1)
    }
    return result
}
