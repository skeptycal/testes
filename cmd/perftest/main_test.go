package main

// import (
// 	"testing"

// 	"github.com/skeptycal/types"
// )

// var any Inccer

// var RetVal interface{}

// type myint int64

// func (i *myint) inc() { *i = *i + 1 }

// type Inccer interface {
// 	inc()
// }

// var inccers = []Inccer{
// 	nil,
// 	new(myint),
// }

// type (
// 	Any = types.Any
// )

// var (

// // benchmarks      = []types.Benchmark{
// // 	types.NewBenchmark("incnAssertNoCheck", incnAssertionNoCheck, []Any{}),
// // 	types.NewBenchmark("incnIntmethod", incnIntmethod, []Any{}),
// // 	types.NewBenchmark("incnInterface", incnInterface, []Any{}),
// // 	types.NewBenchmark("incnSwitch", incnSwitch, []Any{}),
// // 	types.NewBenchmark("incnAssertion", incnAssertion, []Any{}),
// // }

// // bs BenchmarkSet
// )

// func BenchmarkAll(b *testing.B) {

// 	bs = NewBenchmarkSet(b, "(global)", benchmarks)

// 	RetVal.Enable()
// 	bs.Run(b)

// 	bs = NewBenchmarkSet(b, "(local)", benchmarks)

// 	RetVal.Disable()
// 	bs.Run(b)
// }

// // Sample runs (darwin/arm64)
// /*

// //* looping over entire function call
// //* Set global return value (avoiding unwanted compiler optimizations)
// incnAssertNoCheck-8         	 3995966	       310.8 ns/op	      32 B/op	       2 allocs/op
// incnIntmethod-8             	 4037334	       298.8 ns/op	      32 B/op	       2 allocs/op
// incnInterface-8             	 3986919	       301.9 ns/op	      32 B/op	       2 allocs/op
// incnSwitch-8                	 4007119	       298.2 ns/op	      32 B/op	       2 allocs/op
// incnAssertion-8             	 4029920	       297.6 ns/op	      32 B/op	       2 allocs/op

// //* No global return value (compiler optimizations may affect results)
// incnAssertNoCheck-8          	 4034095	       296.2 ns/op	      32 B/op	       2 allocs/op
// incnIntmethod-8              	 4073258	       293.9 ns/op	      32 B/op	       2 allocs/op
// incnInterface-8              	 4055096	       296.9 ns/op	      32 B/op	       2 allocs/op
// incnSwitch-8                 	 4054461	       295.7 ns/op	      32 B/op	       2 allocs/op
// incnAssertion-8              	 4059204	       295.7 ns/op	      32 B/op	       2 allocs/op

// //* tight loops with only necessary code inside
// //* Set global return value (avoiding unwanted compiler optimizations)
// incnAssertNoCheck-8         	147095958	         7.428 ns/op	       0 B/op	       0 allocs/op
// incnIntmethod-8             	177991158	         6.682 ns/op	       0 B/op	       0 allocs/op
// incnInterface-8             	126093896	         9.076 ns/op	       0 B/op	       0 allocs/op
// incnSwitch-8                	140585672	         8.876 ns/op	       0 B/op	       0 allocs/op
// incnAssertion-8             	154314388	         6.760 ns/op	       0 B/op	       0 allocs/op

// //* No global return value (compiler optimizations may affect results)
// incnAssertNoCheck-8          	257806178	         4.676 ns/op	       0 B/op	       0 allocs/op
// incnIntmethod-8              	326270742	         3.687 ns/op	       0 B/op	       0 allocs/op
// incnInterface-8              	236775006	         5.101 ns/op	       0 B/op	       0 allocs/op
// incnSwitch-8                 	257423655	         4.697 ns/op	       0 B/op	       0 allocs/op
// incnAssertion-8              	265160626	         4.611 ns/op	       0 B/op	       0 allocs/op

// */

// func incnAssertionNoCheck(b *testing.B) {
// 	any = new(myint)
// 	// for i := 0; i < b.N; i++ {
// 	any.(*myint).inc()
// 	RetVal.Call(any)
// 	// }
// }
// func incnIntmethod(b *testing.B) {
// 	in := new(myint)
// 	// for i := 0; i < b.N; i++ {
// 	in.inc()
// 	RetVal.Call(in)
// 	// }
// }
// func incnInterface(b *testing.B) {
// 	any = new(myint)
// 	// for i := 0; i < b.N; i++ {
// 	any.inc()
// 	RetVal.Call(any)
// 	// }
// }
// func incnSwitch(b *testing.B) {
// 	any = new(myint)
// 	// for i := 0; i < b.N; i++ {
// 	switch v := any.(type) {
// 	case *myint:
// 		v.inc()
// 	}
// 	RetVal.Call(any)
// 	// }
// }
// func incnAssertion(b *testing.B) {
// 	any = new(myint)
// 	// for i := 0; i < b.N; i++ {
// 	if newint, ok := any.(*myint); ok {
// 		newint.inc()
// 	}
// 	RetVal.Call(any)
// 	// }
// }
