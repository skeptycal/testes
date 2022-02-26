package testes

import (
	"testing"
	"time"
)

var global Any

func BenchmarkAll(b *testing.B) {

	benchmarks := []struct {
		name string
		fn   func(n int) string
		t    time.Duration
	}{
		{"getEncodedString1", getEncodedString1, 0},
		{"getEncodedString2", getEncodedString2, 0},
	}

	const maxCount = 1<<24 - 1

	for _, bb := range benchmarks {
		count := maxCount
		b.Run(bb.name, func(b *testing.B) {
			// for {
			// 	t0 := time.Now()
			for i := 0; i < b.N; i++ {
				global = bb.fn(count)
			}
			// bb.t = time.Since(t0)
			// count *= count
			// if count > maxCount {
			// 	break
			// }
			// }
			b.ReportAllocs()
		})
	}
}
