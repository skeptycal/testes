package testes

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/skeptycal/types"
)

var (
	LimitResult            bool
	DefaultTestResultLimit = 15
)

// ReplacementChar is the recognized unicode replacement
// character for malformed unicode or errors in
// encoding.
//
// It is also found in unicode.ReplacementChar
const ReplacementChar rune = '\uFFFD'

const (
	UPPER    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LOWER    = "abcdefghijklmnopqrstuvwxyz"
	DIGITS   = "0123456789"
	ALPHA    = LOWER + UPPER
	ALPHANUM = ALPHA + DIGITS
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

const (
	maxArgChoices int     = 10
	one           uint64  = 1<<64 - 1
	halfBool      int64   = 1<<63 - 1
	halfHalfBool  int64   = 1<<62 - 1
	ratio         float64 = float64(one) / float64(halfBool)
	halfRatio     float64 = float64(halfBool) / float64(halfHalfBool)
)

func noop(any interface{}) []AnyValue { return nil } // noop function

func boolRatio(n int) float64 {
	m := boolSet(n)
	return float64(m[false]) / float64(m[true])
}

func boolSet(n int) map[bool]int {
	m := make(map[bool]int, 2)
	for i := 0; i < n; i++ {
		m[randBool()]++
	}
	return m
}

func boolSetGroup(n int) []float64 {
	list := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		list = append(list, boolRatio(n))
	}
	return list
}

func stdev(accepted float64, list []float64) float64 {
	w := len(list)
	sum := 0.0

	ac2 := accepted * accepted
	for i := 0; i < w; i++ {

		diff := math.Abs(ac2 - math.Pow(list[i], 2))
		sum += math.Sqrt(diff)
	}

	return sum / float64(w)
}

func RandomString(n int) string {
	sb := strings.Builder{}
	defer sb.Reset()

	for i := 0; i < n; i++ {
		pos := rand.Intn(len(ALPHANUM) - 1)
		sb.WriteByte(ALPHANUM[pos])
	}

	return sb.String()
}

// Contains returns true if the underlying iterable
// sequence (haystack) contains the search term
// (needle) in at least one position.
func Contains(needle Any, haystack []Any) bool {
	for _, x := range haystack {
		if reflect.DeepEqual(needle, x) {
			return true
		}
	}
	return false
}

func TErrorf(t *testing.T, formatString, name string, got, want Any) {
	if formatString == "" {
		formatString = "%v = %v(%T), want %v(%T)"
	}
	t.Errorf(formatString, name, limitTestResultLength(got), got, limitTestResultLength(want), want)
}

func TRunTest(t *testing.T, tt *test) {
	if NewAnyValue(tt.got).IsComparable() && NewAnyValue(tt.want).IsComparable() {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want != tt.wantErr {
				if reflect.DeepEqual(tt.got, tt.want) == tt.wantErr {
					TError(t, tt.name, tt.got, tt.want)
				}
			}
		})
	}
}

func limitTestResultLength(v Any) string {
	s := fmt.Sprintf("%v", v)

	if len(s) > DefaultTestResultLimit && LimitResult {
		return s[:DefaultTestResultLimit-3] + "..."
	}

	return s
}

func TName(testname, funcname, argname Any) string {
	if argname == "" {
		return fmt.Sprintf("%v: %v()", testname, funcname)
	}
	return fmt.Sprintf("%v: %v(%v)", testname, funcname, argname)
}

func typeGuardExclude(needle Any, notAllowed []types.Any) bool {
	return !Contains(needle, notAllowed)
}

func TTypeError(t *testing.T, name string, got, want Any) {
	t.Errorf("%v = %v(%T), want %v(%T)", name, limitTestResultLength(got), got, limitTestResultLength(want), want)
}

func TError(t *testing.T, name string, got, want Any) {
	t.Errorf("%v = %v, want %v", name, limitTestResultLength(got), limitTestResultLength(want))
}
func TTypeRun(t *testing.T, name string, got, want Any, wantErr bool) {
	if NewAnyValue(got).IsComparable() && NewAnyValue(want).IsComparable() {
		t.Run(name, func(t *testing.T) {
			if got != want != wantErr {
				if !reflect.DeepEqual(got, want) {
					TTypeError(t, name, got, want)
				}
			}
		})
	}
}

func TRun(t *testing.T, name string, got, want Any) {
	if NewAnyValue(got).IsComparable() && NewAnyValue(want).IsComparable() {
		t.Run(name, func(t *testing.T) {
			if got != want {
				if !reflect.DeepEqual(got, want) {
					TError(t, name, got, want)
				}
			}
		})
	}
}
