package utils

import "math"

func normL2(vec []float64) float64 {
	var sumSquared float64
	for _, v := range vec {
		sumSquared += v * v
	}
	return math.Sqrt(sumSquared)
}

func Normalize(vec []float64) []float64 {
	norm := normL2(vec)
	for i := range vec {
		vec[i] /= norm
	}
	return vec
}
