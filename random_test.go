package testes

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"testing"
	"time"
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

var smoketest bool = false

func TSmokeTest(t *testing.T, name string) {
	if !smoketest {
		t.Skip("skipping Smoketest - " + name)
	}
}

func Test_randomData(t *testing.T) {
	tests := []struct {
		name    string
		want    Any
		wantErr bool
	}{
		{"smoketest", "smoketest", false},
	}

	TSmokeTest(t, "RandomData")
	LimitResult = true
	for _, tt := range tests {
		for i := 0; i < 1000; i++ {
			name := TName(tt.name, strconv.Itoa(i), "")
			TTypeRun(t, name, RandomData(-1, false), tt.want, tt.wantErr)
		}
	}
}

type sumList struct {
	n    int
	list []int
}

func (s sumList) Get(i int) int      { return s.list[i] }
func (s sumList) Set(i int, v int)   { s.list[i] = v }
func (s sumList) Append(v int)       { s.list = append(s.list, v) }
func (s sumList) Len() int           { return len(s.list) }
func (s sumList) Swap(i, j int)      { s.list[i], s.list[j] = s.list[j], s.list[i] }
func (s sumList) Less(i, j int) bool { return s.list[i] < s.list[j] }
func (s sumList) Reset()             { s.list = s.list[:0] }
func (s sumList) String() string     { return fmt.Sprintf("%v runs of %v : %v\n", s.Len(), s.n, s.list) }

/*
Benchmark_Loops/direct_list_+-8         	1000000000	         0.2279 ns/op	       0 B/op	       0 allocs/op
Benchmark_Loops/slice_+--8              	1000000000	         0.2218 ns/op	       0 B/op	       0 allocs/op
Benchmark_Loops/slice_+_only-8          	1000000000	         0.2229 ns/op	       0 B/op	       0 allocs/op
Benchmark_Loops/map_+--8                	1000000000	         0.2778 ns/op	       0 B/op	       0 allocs/op
Benchmark_Loops/map_true_only-8         	1000000000	         0.2775 ns/op	       0 B/op	       0 allocs/op
Benchmark_Loops/map_false_only-8        	1000000000	         0.2772 ns/op	       0 B/op	       0 allocs/op
Benchmark_Loops/append_+-8              	1000000000	         0.2235 ns/op	       0 B/op	       0 allocs/op
*/
func Benchmark_Loops(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	const n int = 10000
	const nt int = 1000
	sum := 0
	rMap := make(map[bool]int, n)
	trials := sumList{n: n, list: make([]int, nt)}

	b.Run("direct list +", func(b *testing.B) {
		for j := 0; j < nt; j++ {
			for i := 0; i < n; i++ {
				if randBool() {
					trials.list[j]++
				}
			}
		}
	})

	global = trials
	// fmt.Println(trials)
	trials.Reset()

	b.Run("slice +-", func(b *testing.B) {
		for j := 0; j < nt; j++ {
			sum = 0
			for i := 0; i < n; i++ {
				if randBool() {
					sum++
				} else {
					sum--
				}
			}
			trials.Set(j, sum)
		}
	})

	global = trials
	// fmt.Println(trials)
	trials.Reset()

	b.Run("slice + only", func(b *testing.B) {
		for j := 0; j < nt; j++ {
			sum = 0
			for i := 0; i < n; i++ {
				if randBool() {
					sum++
				}
			}
			trials.Set(j, sum)
		}
	})

	global = trials
	// fmt.Println(trials)
	trials.Reset()

	b.Run("map +-", func(b *testing.B) {
		for j := 0; j < nt; j++ {
			rMap[true] = 0
			rMap[false] = 0
			for i := 0; i < n; i++ {
				rMap[randBool()]++
			}
			trials.Set(j, rMap[true]-rMap[false])

		}
	})
	rMap[true] = 0
	rMap[false] = 0

	global = trials
	// fmt.Println(trials)
	trials.Reset()

	b.Run("map true only", func(b *testing.B) {
		for j := 0; j < nt; j++ {
			rMap[true] = 0
			for i := 0; i < n; i++ {
				rMap[randBool()]++
			}
			trials.Set(j, rMap[true])
		}
	})

	rMap[true] = 0
	rMap[false] = 0

	global = trials
	// fmt.Println(trials)
	trials.Reset()

	b.Run("map false only", func(b *testing.B) {
		for j := 0; j < nt; j++ {
			rMap[false] = 0
			for i := 0; i < n; i++ {
				rMap[randBool()]++
			}
			trials.Set(j, rMap[false])
		}
	})

	global = trials
	// fmt.Println(trials)
	trials.Reset()

	b.Run("append +", func(b *testing.B) {
		for j := 0; j < nt; j++ {
			sum = 0
			for i := 0; i < n; i++ {
				if randBool() {
					sum++
				}
			}
			trials.Append(sum)
		}
	})

	global = trials
	// fmt.Println(trials)
	trials.Reset()

}

func Test_randBool(t *testing.T) {
	TSmokeTest(t, "randBool")

	n := 1000
	nt := 10000
	sum := 0
	trials := make([]int, nt)
	rMap := make(map[bool]int, n)
	for j := 0; j < nt; j++ {
		for i := 0; i < n; i++ {
			rMap[randBool()]++
			if randBool() {
				sum++
			} else {
				sum--
			}
		}
		trials[j] = sum
		// t.Errorf("randBool ratio (n=%v) = , %v want %v", n, sum, 0)
	}

	mean := Mean(trials)
	sd := StDev(trials)

	t.Run("randBool", func(t *testing.T) {
		if sum != 0 {
			t.Errorf("randBool ratio (%v trials of n=%v) = mean (sd) %v (%v) want %v", nt, n, mean, sd, 0)
		}
	})

}
