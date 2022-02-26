package testes

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

const arune = 'A'

type (
	KindMap map[string]int

	StatMap interface {
		Keys() []string
		Values() []int
		Min() int
		Max() int
		Mean() float64
		StDev() float64

		String() string
	}

	dataPoint struct {
		t0   time.Time
		t1   time.Time
		fn   func() Any
		data Any
	}

	DataPoint interface {
		Start()
		Stop()
		Collect()
	}
)

func (m KindMap) Keys() []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (m KindMap) Values() []int {
	values := make([]int, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	sort.Ints(values)
	return values
}

func (m KindMap) Min() int {
	min := 1<<63 - 1
	for _, v := range m {
		if v < min {
			min = v
		}
	}
	return min
}

func (m KindMap) Max() int {
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

func (m KindMap) Mean() float64 {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return float64(sum) / float64(len(m))
}

func (m KindMap) String() string {
	sb := strings.Builder{}
	defer sb.Reset()

	for k, v := range m {
		sb.WriteString(fmt.Sprintf("%v = %v\n", k, v))
	}
	return sb.String()
}

func dev(x, m float64) float64 {
	return (x - m) * (x - m)
}

func (m KindMap) StDev() float64 {
	var sum, x, mean float64
	mean = m.Mean()

	for _, v := range m {
		x = float64(v)
		sum += (x - mean) * (x - mean)
		// sum += dev(x, mean)
	}
	return math.Sqrt(sum / float64(len(m)))
}

func Mean(list []int) float64 {
	var sum int
	for _, v := range list {
		sum += v
	}
	return float64(sum) / float64(len(list))
}

func StDev(list []int) float64 {
	mean := Mean(list)
	var sum float64
	for _, v := range list {
		x := float64(v)
		sum += (x - mean) * (x - mean)
	}
	return math.Sqrt(sum / float64(len(list)))

}

func Less(i, j int) bool {
	return i < j
}

func (d *dataPoint) Start() {
	d.t0 = time.Now()
}

func (d *dataPoint) Stop() {
	d.t1 = time.Now()
}

func (d *dataPoint) Collect() {
	d.Start()
	d.data = d.fn()
	d.Stop()
}

func GetData() DataPoint {
	d := new(dataPoint)
	d.Collect()
	return d
}

func getEncodedString1(n int) string {
	b := []byte("")
	for i := 0; i < n; i++ {
		b = append(b, byte(RandomKind(true))+65)
	}
	return string(b)
}

func getEncodedString2(n int) string {
	sb := strings.Builder{}
	defer sb.Reset()
	for i := 0; i < n; i++ {
		sb.WriteByte(byte(RandomKind(true) + 65))
	}
	return sb.String()
}

func GetEncodedString(n int) string {
	return getEncodedString1(n)
}
