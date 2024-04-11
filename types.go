package comp

type SignedIntegerType interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}
type UnsignedIntegerType interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~uintptr
}
type IntegerType interface {
	SignedIntegerType | UnsignedIntegerType
}
type Float interface {
	~float32 | ~float64
}
type Complex interface {
	~complex64 | ~complex128
}
type FloatingPoint interface {
	Float | Complex
}
type Real interface {
	IntegerType | Float
}
type Arithmetic interface {
	IntegerType | Float | Complex
}
type Ordered interface {
	IntegerType | Float | ~string
}
type BooleanTestible interface {
	~bool | Arithmetic | ~string
}

func Abs[A Real](a A) A {
	return Cond(a < 0, -a, a).(A)
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
	default:
		return a != nil
	}
}

// operator ?:
func Cond(cond any, t any, f any) any {
	if Bool(cond) {
		return t
	} else {
		return f
	}
}

// operator ??
func Null(v any, w any) any {
	if v != nil {
		return v
	} else {
		return w
	}
}

func Value(v any, err error) any {
	if err == nil {
		return v
	} else {
		return err
	}
}
