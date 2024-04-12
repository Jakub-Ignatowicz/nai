package utils

import (
	"errors"
)

func Perceptron(x []float64, w []float64) (float64, error) {
	if len(x) != len(w) {
		return 0, errors.New("Entry vector and weights vector are not the same lenghts")
	}

	var sum float64 = 0
	for i := 0; i < len(x); i++ {
		sum += (x[i] * w[i])
	}

	return sum, nil
}
func TeachPerceptron(expectedValue float64, x []float64, w []float64, alpha float64) ([]float64, error) {
	perceptron, err := Perceptron(x, w)
	if err != nil {
		return nil, err
	}

	error := expectedValue - perceptron

	for i := 0; i < len(x); i++ {
		w[i] += alpha * error * x[i]
	}

	return w, nil
}
