package testes

import (
	"runtime"
	"strconv"
	"testing"
)

func Benchmark_makeRandomNumbers(b *testing.B) {

	// TODO: likely not needed but available for benchmarking
	// Since Go 1.5, GOMAXPROCS defaults to the number of
	// CPU cores available, so no need to set that (although
	// it does no harm).
	// Reference: https://stackoverflow.com/a/41632900
	numThreads := runtime.NumCPU()
	runtime.GOMAXPROCS(numThreads)

	numIntsToGenerate = 1000

	ch := make(chan int, 1000)

	singleThreadIntSlice := make([]int, 0, numIntsToGenerate)
	multiThreadIntSlice := make([]int, 0, numIntsToGenerate)

	b.RunParallel(func(pb *testing.PB) {
		for i := 0; i < numIntsToGenerate; i++ {
			singleThreadIntSlice[i] = <-ch
		}
	})

	b.RunParallel(func(pb *testing.PB) {
		for i := 0; i < numIntsToGenerate; i++ {
			multiThreadIntSlice[i] = <-ch
		}
	})

}

func Test_makeRandomNumbers(t *testing.T) {
	tests := []struct {
		name    string
		numInts int
		ch      chan int
	}{
		// TODO: Add test cases.
		{"10", 10, make(chan int, 10)},
		{"100", 100, make(chan int, 100)},
		{"1000", 1000, make(chan int, 1000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			makeRandomNumbers(tt.numInts, tt.ch)
			TRun(t, "channel filled", len(tt.ch), tt.numInts)
		})
	}
}

func Test_randomData(t *testing.T) {
	tests := []struct {
		name    string
		want    Any
		wantErr bool
	}{
		// TODO: Add test cases.
		{"smoketest", "smoketest", false},
	}

	LimitResult = true
	for _, tt := range tests {
		for i := 0; i < 1000; i++ {
			name := TName(tt.name, strconv.Itoa(i), "")
			TTypeRun(t, name, RandomData(-1, false), tt.want, tt.wantErr)
		}
	}
}
