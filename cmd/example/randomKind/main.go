package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/skeptycal/testes"
)

const arune = 'A'

type kindMap map[string]int

func (m kindMap) Keys() []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (m kindMap) Values() []int {
	values := make([]int, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	sort.Ints(values)
	return values
}

func (m kindMap) Min() int {
	min := 1<<63 - 1
	for _, v := range m {
		if v < min {
			min = v
		}
	}
	return min
}

func (m kindMap) Max() int {
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

func (m kindMap) Mean() float64 {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return float64(sum) / float64(len(m))
}

func (m kindMap) String() string {
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

func (m kindMap) StDev() float64 {
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

type (
	Any = interface{}

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
		b = append(b, byte(testes.RandomKind(true))+65)
	}
	return string(b)
}

func getEncodedString2(n int) string {
	sb := strings.Builder{}
	defer sb.Reset()
	for i := 0; i < n; i++ {
		sb.WriteByte(byte(testes.RandomKind(true) + 65))
	}
	return sb.String()
}

func GetEncodedString(n int) string {
	return getEncodedString1(n)
}

func main() {
	const n = 1<<24 - 1

	k := make(kindMap, n)

	for i := 0; i < n; i++ {
		kn := testes.RandomKind(true)
		k[kn.String()]++
	}

	GetEncodedString(n)

	// fmt.Println(string(b))

	// list := k.Values()
	// fmt.Println(list)
	// fmt.Printf("Mean(list): %v\n", Mean(list))
	// fmt.Printf("StDev(list): %v\n", StDev(list))

	fmt.Println(k)

	mean := k.Mean()
	sd := k.StDev()
	min := k.Min()
	max := k.Max()

	fmt.Printf("map.Min(): %v\n", min)
	fmt.Printf("map.Max(): %v\n", max)
	fmt.Printf("map.Mean(): %v\n", mean)
	fmt.Printf("map.StDev(): %v\n", sd)

	times := make([]DataPoint, 0, 16)

	for i := 0; i < n; i++ {
		times = append(times, GetData())
	}

	fmt.Println("Times: ", times)

}
