package popcount4

import "testing"

func BenchmarkPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCount(uint64(5))
    }
}
