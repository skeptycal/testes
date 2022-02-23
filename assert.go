package testes

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/skeptycal/types"
)

var NewAnyValue = types.NewAnyValue

const (
	assertEqual     = "AssertEqual(%v): got %v, want %v"
	assertNotEqual  = "AssertNotEqual(%v): got %v, want %v"
	assertDeepEqual = "AssertDeepEqual(%v): got %v, want %v"
	assertSameType  = "AssertSameType(%v): got %v, want %v"
	assertSameKind  = "AssertSameKind(%v): got %v, want %v"
	assertSameFunc  = "AssertSameFunc(%v): got %v, want %v"
)

func IsKindEqual(got, want Any) bool {
	if gv, ok := got.(reflect.Kind); ok {
		if wv, ok := want.(reflect.Kind); ok {
			return gv == wv
		}
	}
	return false
}

func AssertKindEqual(t *testing.T, name string, got, want Any) bool {
	if IsKindEqual(got, want) {
		return true
	}
	TErrorf(t, assertSameKind, name, got, want)
	return false
}

func GetFuncResult(t *testing.T, name string, fn Any, args ...reflect.Value) ([]reflect.Value, error) {
	if fn == nil {
		return nil, fmt.Errorf("fn must be provided: %v", fn)
	}

	f := NewAnyValue(fn)

	if !IsKindEqual(f.Kind(), reflect.Func) {
		return nil, fmt.Errorf("fn must be a function: %v(%v)", fn, f.Kind())
	}

	return nil, nil
}

// AssertSameFunc returns true if got and want are
// both functions that return the same value when
// called with args... as input.
func AssertSameFunc(t *testing.T, name string, got, want Any, args ...reflect.Value) bool {

	g := NewAnyValue(got)

	// If got is a pointer, run again with object
	if IsKindEqual(g.Kind(), reflect.Ptr) {
		return AssertSameFunc(t, name, reflect.Indirect(g.ValueOf()), want)
	}

	// If want is a pointer, run again with object
	w := NewAnyValue(want)
	if w.Kind() == reflect.Ptr {
		return AssertSameFunc(t, name, got, reflect.Indirect(w.ValueOf()))
	}
	if g.Kind() != reflect.Func {
		TErrorf(t, assertSameFunc, name, g.Kind(), reflect.Func)
		return false
	}

	if g.Kind() != reflect.Func {
		TErrorf(t, assertSameFunc, name, g.Kind(), reflect.Func)
		return false
	}

	gf := g.ValueOf().Call(args)
	wf := w.ValueOf().Call(args)

	return AssertEqual(t, name, gf, wf)

}

func AssertEqual(t *testing.T, name string, got, want Any) bool {
	if got == want {
		return true
	}
	t.Errorf(assertEqual, name, got, want)
	return false
}

func AssertNotEqual(t *testing.T, name string, got, want Any) bool {
	if got != want {
		return true
	}
	t.Errorf(assertNotEqual, name, got, want)
	return false
}

func AssertDeepEqual(t *testing.T, name string, got, want Any) bool {
	if reflect.DeepEqual(got, want) {
		return true
	}
	t.Errorf(assertDeepEqual, name, got, want)
	return false
}

func AssertSameType(t *testing.T, name string, got, want Any) bool {
	g := NewAnyValue(got).TypeOf()
	w := NewAnyValue(want).TypeOf()

	if g == w {
		return true
	}
	t.Errorf(assertSameType, name, g, w)
	return false
}

func AssertSameKind(t *testing.T, name string, got, want Any) bool {
	g := NewAnyValue(got).Kind()
	w := NewAnyValue(want).Kind()

	if g == w {
		return true
	}
	t.Errorf(assertSameKind, name, g, w)
	return false
}
