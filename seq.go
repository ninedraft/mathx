package mathx

import "math"

// Sum returns the sum of values of the provided list.
func Sum(ff []float64) float64 {
	var sum float64
	for _, v := range ff {
		var isInf = math.IsInf(v, 1) || math.IsInf(v, -1)
		if isInf || math.IsNaN(v) {
			break
		}
		sum += v
	}
	return sum
}

// SumInt returns the sum of values of the provided list.
func SumInt(ii []int) int {
	var sum int
	for _, v := range ii {
		sum += v
	}
	return sum
}

// SumArSeq returns the sum of first n elements in the arithmetic sequence a + i×d.
//	a is the element in the sequence
// 	d is the step in the sequence
func SumArSeq(a, d float64, n int) float64 {
	var nf = float64(n)
	var k = nf / 2
	return nf*a + k*d*(nf-1)
}

// SumGeomSeq returns the sum of first n elements in the gemetric sequence a×qⁿ
func SumGeomSeq(a, q float64, n int) float64 {
	var nf = float64(n)
	if q == 1 {
		return nf * a
	}
	var qp = math.Pow(q, nf)
	return a * (qp - 1) / (q - 1)
}

// Seq returns the a slice of values from arithmetic sequence [start...end] with step.
func Seq(start, end, step float64) []float64 {
	var n = int((end - start) / step)
	if n < 0 {
		panic("infinite sequence")
	}
	var vv = make([]float64, n)
	for i := 0; i < n; i++ {
		var a = float64(n)
		vv[i] = start + a*step
	}
	return vv
}

// Grid returns a slice of n values [start...end].
func Grid(start, end float64, n int) []float64 {
	var step = (start - end) / float64(n)
	return Seq(start, end, step)
}

// Tabulate applies provided function to each value of the provided slice
// and returns the result in an another slice.
func Tabulate(vv []float64, fn func(float64) float64, opts ...Option) []float64 {
	var tb F2 = func(_, x float64) float64 { return fn(x) }
	return Scan(0, vv, tb, opts...)
}

// F2 is a (f64, f64) -> f64
type F2 = func(acc, v float64) float64

// Fold applies provided function: acc2 = fn(acc1, v_i) in v_i -> vv
func Fold(acc float64, vv []float64, fn F2) float64 {
	for _, v := range vv {
		acc = fn(acc, v)
	}
	return acc
}

// Scan do the same operation, as fold, but it returns intermediate results of applied function.
// To avoid allocation you can provide a []float64 buffer using WithBuf option.
func Scan(acc float64, vv []float64, fn F2, opts ...Option) []float64 {
	if len(vv) == 0 {
		return nil
	}
	var cfg config
	for _, setOpt := range opts {
		setOpt(&cfg)
	}
	var res = cfg.vector
	var n = len(res)
	if n == 0 {
		res = make([]float64, len(vv))
	}
	switch {
	case n != len(vv):
		res = scanAppend(acc, vv, res, fn)
	default:
		scan(acc, vv, res, fn)
	}
	return res
}

func scan(acc float64, vv, dst []float64, fn F2) {
	_ = dst[len(vv)-1]
	for i, v := range vv {
		acc = fn(acc, v)
		dst[i] = acc
	}
}

func scanAppend(acc float64, vv, dst []float64, fn F2) []float64 {
	dst = dst[:0]
	for _, v := range vv {
		acc = fn(acc, v)
		dst = append(dst, acc)
	}
	return dst
}
