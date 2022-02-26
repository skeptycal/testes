package main

import (
	"fmt"

	"github.com/skeptycal/testes"
	"github.com/skeptycal/types"
)

var global types.Any

func main() {
	const n = 1<<24 - 1

	k := make(testes.KindMap, n)

	for i := 0; i < n; i++ {
		kn := testes.RandomKind(true)
		k[kn.String()]++
	}

	s := testes.GetEncodedString(n)

	fmt.Println(s)

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

	// times := make([]testes.DataPoint, 0, 16)

	// for i := 0; i < n; i++ {
	// 	times = append(times, testes.GetData())
	// }

	// fmt.Println("Times: ", times)

}
