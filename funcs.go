package comp

import (
	"math"
	"math/cmplx"
)

// Absolute valueReal
func Abs(v any) any {
	switch a := v.(type) {
	case complex64:
		return complex64(cmplx.Abs(complex128(a)))
	case complex128:
		return cmplx.Abs(a)
	case int:
		return Cond(a < 0, -a, a)
	case int8:
		return Cond(a < 0, -a, a)
	case int16:
		return Cond(a < 0, -a, a)
	case int32:
		return Cond(a < 0, -a, a)
	case int64:
		return Cond(a < 0, -a, a)
	case float32:
		return float32(math.Abs(float64(a)))
	case float64:
		return math.Abs(a)
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return a
	default:
		return v //unnecessary to panic
	}
}

func Bool(b any) bool {
	switch a := b.(type) {
	case bool:
		return a
	case string:
		return a != ""
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64, complex64, complex128:
		return a != 0
	case interface{ Bool() bool }:
		return a.Bool()
	case unsafe.Pointer:
		return uintptr(a) != 0
	default:
		return a != nil
	}
}

// Operator ?:
func Cond(cond, t, f any) any {
	if Bool(cond) {
		return t
	} else {
		return f
	}
}

// Operator ??
func Nil(v any, w any) any {
	if v != nil {
		return v
	} else {
		return w
	}
}

// Obtain the value if no error, or the error if one
func Value(v any, err error) any {
	if err == nil {
		return v
	} else {
		return err
	}
}
