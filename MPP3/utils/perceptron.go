package utils

import (
	"errors"
	"math"
)

func Perceptron(x []float64, w []float64) (float64, error) {
	if len(x) != len(w) {
		return 0, errors.New("Entry vector and weights vector are not the same lenghts")
	}
	normalizedX := normalize(x)
	normalizedW := normalize(w)
	var sum float64 = 0
	for i := 0; i < len(x); i++ {
		sum += (normalizedX[i] * normalizedW[i])
	}

	return sum, nil
}

func TeachPerceptron(expectedValue float64, x []float64, w []float64, alpha float64) ([]float64, error) {
	per, err := Perceptron(x, w)
	if err != nil {
		return nil, err
	}

	correctionRate := expectedValue - per

	for i := 0; i < len(x); i++ {
		w[i] += x[i] * alpha * correctionRate
	}

	return w, nil
}

func normL2(vec []float64) float64 {
	var sumSquared float64
	for _, v := range vec {
		sumSquared += v * v
	}
	return math.Sqrt(sumSquared)
}

func normalize(vec []float64) []float64 {
	norm := normL2(vec)
	for i := range vec {
		vec[i] /= norm
	}
	return vec
}
