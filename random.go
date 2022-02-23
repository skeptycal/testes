package testes

import (
	"math/cmplx"
	"math/rand"
	"reflect"
	"time"
	"unsafe"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// 1 million is enough now:
var numIntsToGenerate = 1000 * 1000

func makeRandomNumbersBench(numInts int, ch chan int) {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	for i := 0; i < numInts; i++ {

		// TODO: This step is only for benchmarking results:
		// Kill time, do some processing:
		for j := 0; j < 1000; j++ {
			generator.Intn(numInts * 100)
		}

		// and now return a single random number
		ch <- generator.Intn(numInts * 100)
	}
}

func makeRandomNumbers(numInts int, ch chan int) {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	for i := 0; i < numInts; i++ {
		ch <- generator.Intn(numInts * 100)
	}
}

func randomKind(useInvalid bool) reflect.Kind {
	if useInvalid {
		return reflect.Kind(rand.Intn(26))
	}
	return reflect.Kind(rand.Intn(25) + 1)
}

func randBool() bool {
	return rand.Int63() >= halfHalfBool
}

const (
	b8  = 1<<8 - 1
	b16 = 1<<16 - 1
	b32 = 1<<32 - 1
	b64 = 1<<64 - 1
	u8  = b64 >> 56
	u16 = b64 >> 48
)

func Conj(c complex128) complex128 {
	return cmplx.Conj(c)
}

func rng(n int) int { return rand.Intn(n) }

func RandomData(useInvalid bool) Any {
	switch k := randomKind(useInvalid); k {
	case 0: // Invalid
		if useInvalid {
			return k
		}
		return RandomData(false)
	case 1: // Bool
		return randBool()
	case 2: // Int
		return rand.Int()
	case 3: // Int8
		return int8(rand.Intn(b8))
	case 4: // Int16
		return int16(rand.Intn(b16))
	case 5: // Int32
		return int32(rand.Uint32())
	case 6: // Int64
		return int64(rand.Uint64())
	case 7: // Uint
		return uint(rand.Uint64())
	case 8: // Uint8
		return rand.Uint64() >> 56
	case 9: // Uint16
		return rand.Uint64() >> 48
	case 10: // Uint32
		return rand.Uint32()
	case 11: // Uint64
		return rand.Uint64()
	case 12: // Uintptr
	case 13: // Float32
		return rand.Float32()
	case 14: // Float64
		return rand.Float64()
	case 15: // Complex64
		return complex(rand.Float32(), rand.Float32())
	case 16: // Complex128
		return complex(rand.Float64(), rand.Float64())
	case 17: // Array
		a := [16]int{}
		for i := range a {
			a[i] = rng(10) + 5
		}
		return a
	case 18: // Chan
		return make(chan int, rng(10)+5)
	case 19: // Func
		return func() string { return "fake function" }
	case 20: // Interface
		return NewAnyValue(rng(42))
	case 21: // Map
		m := make(map[int]bool, rng(10)+5)
		for i := 0; i < len(m); i++ {
			m[i] = randBool()
		}
		return m
	case 22: // Ptr
		value := int(rand.Uint64())
		return &value
	case 23: // Slice
		s := make([]bool, 0, rng(10)+5)
		for i := 0; i < len(s); i++ {
			s[i] = randBool()
		}
		return s
	case 24: // String
		return RandomString(rng(10) + 5)
	case 25: // Struct
		return struct {
			name  string
			value int
		}{RandomString(rng(10) + 5), rng(10) + 5}
	case 26: // UnsafePointer
		value := int(rand.Uint64())
		return unsafe.Pointer(&value)
	}
	return nil
}
