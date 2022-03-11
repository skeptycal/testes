package testes

import (
	"fmt"
	"reflect"
	"testing"

	"google.golang.org/appengine/log"
)

/*

Automated testing and benchmarking:

tests:

gather function list
gather input data list

call tests and report

benchmarks:

gather function list
gather input data list


loop through b.N {
	loop through all data sets
}

*/

func TBbN(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}

// NoTestResult is a slice containing only one value, the
// zero value of reflect.Value, which is an invalid result.
var NoValues TBvalues = []reflect.Value{reflect.Value{}}

// NoValue is the zero value of reflect.Value, which is an
// invalid result. It specifically means that the test
// result has not yet been determined.
var NoValue reflect.Value = reflect.Value{}

func NewTest(t *testing.T, name string, fn anyFunc, args, want []Any, wantErr bool) (Test, error) {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		return nil, fmt.Errorf("fn must be a function: %v(%T)", fn, fn)
	}

	return &testStruct{
		name:    name,
		fn:      v,
		in:      TBwrapInputValues(args),
		want:    TBwrapInputValues(want),
		got:     TBvalues{NoValue},
		wantErr: wantErr,
	}, nil

}

type (
	anyFunc = interface{}

	TBvalues = []reflect.Value

	testStruct struct {
		name    string        // descriptive name
		fn      reflect.Value // must be a function
		in      TBvalues      // slice of input arguments
		want    TBvalues      // slice of return values
		got     TBvalues      // actual function result (jit)
		wantErr bool          // expecting an error?
	}

	TestSet interface {
	}

	Test interface {

		// Name returns the name of the test.
		Name() string

		// Args returns the function input arguments.
		Args() TBvalues

		// Want returns the expected return value.
		Want() TBvalues

		// WantErr returns true if an error is expected
		// from an AssertTrue test (got == want).
		WantErr() bool

		// Got returns the cached function result. If the
		// function has not yet been called, it is called
		// and the value is cached and returned.
		Got() TBvalues

		// Run calls the the function (jit) if it has not
		// been called before and caches the result.
		Run()

		// Call specifically calls the function and returns
		// the result of the call whether the result is cached
		// or not.
		Call(in TBvalues) TBvalues
	}
)

func (f *testStruct) Name() string   { return f.name }
func (f *testStruct) Args() TBvalues { return f.in }
func (f *testStruct) Want() TBvalues { return f.want }
func (f *testStruct) WantErr() bool  { return f.wantErr }
func (f *testStruct) Run()           { f.got = f.Call() }
func (f *testStruct) Got() TBvalues {
	if f.got[1] == NoValue {
		f.Run()
	}
	return f.got
}

func (f *testStruct) Call() TBvalues {
	defer func() {
		if r := recover(); r != nil {
			Log.Errorf("TestBench: recovered from panic while calling %v: %v", f.name, r)
		}
	}()
	return f.fn.Call(f.in)
}

///////////////////

func TBmakeArgList(tests ...Test) []TBvalues {
	n := len(tests)
	argList := make([]TBvalues, 0, n)
	for j := 0; j < n; j++ {
		argList = append(argList, TBwrapInputValues(tests[j]))
	}
	return argList
}

func TBwrapInputValues(args ...interface{}) TBvalues {
	if len(args) == 0 {
		return TBvalues{}
	}

	v := make(TBvalues, 0, len(args))
	for _, a := range args {
		v = append(v, reflect.ValueOf(a))
	}
	return v
}

// TBloop returns the output output values from a call
// to the function fn using argList as a series of inputs.
// ArgList is a slice of inputs; fn is called once for
// each argument in argList and is results are returned in
// a similar slice of return values.
var TBloop = func(fn Any, in []TBvalues) ([]TBvalues, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("TestBench: recovered from panic: %v (%v)", r, fn)
		}
	}()

	a := make([]TBvalues, 0, len(in))

	for i := 0; i < len(in); i++ {
		a = append(a, v.Call(in[i]))
	}
	return a, nil
}
