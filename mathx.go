package mathx

import (
	"unsafe"

	"golang.org/x/exp/constraints"
)

type Number interface {
	Real | Complex
}

type Real interface {
	Integer | Float
}

type Signed interface {
	constraints.Signed | Float
}

type Complex = constraints.Complex

type Integer = constraints.Integer

type Float = constraints.Float

func IsNaN[N Number](x N) bool {
	return x != x
}

func Max[R Real](x, y R) R {
	if IsNaN(x) {
		return x
	}

	if IsNaN(y) {
		return y
	}

	if x == 0 && x == y {
		if Signbit(x) {
			return y
		}
		return x
	}

	if x > y {
		return x
	}

	return y
}

func Min[R Real](x, y R) R {
	if IsNaN(x) {
		return x
	}

	if IsNaN(y) {
		return y
	}

	if x == 0 && x == y {
		if Signbit(x) {
			return x
		}
		return y
	}

	if x < y {
		return x
	}

	return y
}

func Sign[S Signed](x S) S {
	if IsNaN(x) {
		return x
	}

	if x == 0 {
		return 0
	}

	if Signbit(x) {
		return -1
	}

	return 1
}

func Signbit[R Real](x R) bool {
	n := unsafe.Sizeof(x)
	ptr := (*byte)(unsafe.Pointer(&x))
	slice := unsafe.Slice(ptr, n)

	return slice[n-1]&(1<<7) != 0
}
