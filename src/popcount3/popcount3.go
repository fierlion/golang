package popcount3

func PopCount(x uint64) int {
    result := 0
    for i := 0; i < 64; i++ {
        if ((x & 1) == 1) {
            result += 1
        }
        x = x>>1
    }
    return result
}
