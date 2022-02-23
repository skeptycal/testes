package testes

import (
	"testing"
)

type (
	test struct {
		t       *testing.T
		name    string
		in      Any
		got     Any
		want    Any
		wantErr bool
	}

	assert struct {
		tt *test
	}

	// Assert implements the Tester interface. It is
	// used for boolean only challenges. In addition
	// to working seamlessly with the standard library
	// testing package, it can return the bool
	// result for use in alternate data collection
	// or CI software. Result() automatically calls
	// Run()
	Assert interface {
		Tester
		Result() bool
	}

	// Random implements Tester and  creates a random
	// test that can be used to generate many varied
	// tests automatically.
	// After each use, Regenerate() can be called to
	// generate a new test.
	// Regenerate() automatically calls Run().
	Random interface {
		Tester
		Regenerate()
	}

	// Custom implements Tester and can be used to
	// hook into existing software by passing in
	// the various test arguments with Hook().
	// Calling Hook() also calls Run() automaticaly.
	Custom interface {
		Tester
		Hook(name string, got, want Any, wantErr bool)
	}

	testSet struct {
		name string
		t    *testing.T
		list []Tester
	}

	// Tester implements the Run method of an automated
	// test suite. It may be implemented by traditional
	// tests, asserts, random inputs, custom code, or
	// sets of tests.
	Tester interface {
		Run()
	}
)

func NewTestSet(t *testing.T, name string, list []Tester) Tester {
	return &testSet{
		t:    t,
		name: name,
		list: list,
	}
}

// func NewAssert() Assert {

// }

func NewTest(t *testing.T, name string, in, got, want Any, wantErr bool) Tester {
	return &test{
		t:       t,
		name:    name,
		in:      in,
		got:     got,
		want:    want,
		wantErr: wantErr,
	}
}

// Run runs all tests in the set.
func (ts *testSet) Run() {
	for _, tt := range ts.list {
		tt.Run()
	}
}

// Run runs the individual test
func (tt *test) Run() {
	TRunTest(tt.t, tt)
}
