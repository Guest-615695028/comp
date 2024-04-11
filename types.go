package comp
type SignedInt interface{ ~int8 | ~int16 | ~int32 | ~int64 | ~int }
type UnsignedInt interface{ ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~uintptr }
type Int interface{ SignedInt | UnsignedInt }
type Float interface{ ~float32 | ~float64 }
type Complex interface{ ~complex64 | ~complex128 }
type FloatingPoint interface{ Float | Complex }
type Real interface{ Int | Float }
type Arithmetic interface{ Int | Float | Complex }
type Ordered interface{ Int | Float | ~string }
type Booler interface{ Bool() bool }
type BooleanTestible interface{ ~bool | ~string | Arithmetic }
