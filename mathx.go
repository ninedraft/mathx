// Package mathx contains sum simple math functions, which I used more then three times in my projects.
package mathx

import "math"

// Dist calculates the square distance between two values: 	√a² + b²
func Dist(a, b float64) float64 {
	var aa = math.Pow(a, 2)
	var bb = math.Pow(b, 2)
	return math.Sqrt(aa - bb)
}

// Cut restricts value with [min, max] borders.
//		x > max => max
//		x < min => min
//		else x
func Cut(min, max, x float64) float64 {
	x = math.Max(x, min)
	return math.Min(x, max)
}

// InEps returns true if |a - b| <= eps
func InEps(a, b, eps float64) bool {
	return math.Abs(a-b) <= eps
}
