package mathx

import "math"

// Min returns the minimum value in the list of values [a, rest...]
// If rest is empty, the result is a.
func Min(a float64, rest ...float64) float64 {
	var min = a
	for _, v := range rest {
		switch {
		case math.IsNaN(v), math.IsInf(v, -1):
			return v
		default:
			min = math.Min(min, v)
		}
	}
	return min
}

// MinIndex returns the minimum value in the list of values.
// Returns index of first -Inf value.
// Returns -1 if the list is empty or contains a NaN value.
// If there several entries of the minimum in the list,
// then returns the index of the last occurrence.
func MinIndex(ff []float64) int {
	var index = -1
	if len(ff) == 0 {
		return index
	}
	var min = ff[0]
	for i, v := range ff[1:] {
		switch {
		case math.IsInf(v, -1):
			return i
		case math.IsNaN(v):
			return -1
		case v == math.Min(v, min):
			index = i
			min = v
		}
	}
	return index
}

// Max returns the maximum value in the list of values [a, rest...]
// If rest is empty, the result is a.
func Max(a float64, rest ...float64) float64 {
	var max = a
	for _, v := range rest {
		switch {
		case math.IsNaN(v), math.IsInf(v, 1):
			return v
		default:
			max = math.Max(max, v)
		}
	}
	return max
}

// MaxIndex returns the maximum value in the list of values.
// Returns index of first +Inf value.
// Returns -1 if the list is empty or contains a NaN value.
// If there several entries of the maximum in the list,
// then returns the index of the last occurrence.
func MaxIndex(ff []float64) int {
	var index = -1
	if len(ff) == 0 {
		return index
	}
	var max = ff[0]
	for i, v := range ff[1:] {
		switch {
		case math.IsInf(v, 1), math.IsNaN(v):
			return -1
		case v == math.Max(v, max):
			index = i
			max = v
		}
	}
	return index
}

// MinInt returns the smallest integer from the list of [a, rest...].
func MinInt(a int, rest ...int) int {
	var min = a
	for _, v := range rest {
		min = minInt(min, v)
	}
	return min
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt returns the smallest integer from the list of [a, rest...].
func MaxInt(a int, rest ...int) int {
	var max = a
	for _, v := range rest {
		max = maxInt(max, v)
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
