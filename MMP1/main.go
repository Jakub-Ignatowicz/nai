package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Flower struct {
	name         string
	measurements []float32
}

func (f1 *Flower) Distance(f2 *Flower) float32 {
	var maxLenght int
	if len(f1.measurements) > len(f2.measurements) {
		maxLenght = len(f1.measurements)
	} else {
		maxLenght = len(f2.measurements)
	}

	distance := 0.0

	for i := 0; i < maxLenght; i++ {
		var measurement1, measurement2 float32
		if i < len(f1.measurements) {
			measurement1 = f1.measurements[i]
		}
		if i < len(f2.measurements) {
			measurement2 = f2.measurements[i]
		}

		diff := measurement1 - measurement2
		distance += math.Pow(float64(diff), 2)
	}

	return float32(math.Sqrt(distance))
}

func main() {
	flowers, err := flowerDataReader("./data/iris_training.txt")

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range flowers {
		fmt.Println(v)
	}
}

func flowerDataReader(fileName string) ([]Flower, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, errors.New("Unable to open a file: " + err.Error())
	}
	defer file.Close()

	var flowers = make([]Flower, 0)

	counter := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		counter++
		line := scanner.Text()
		flowerData := strings.Fields(line)
		if len(flowerData) < 2 {
			return nil, errors.New("Incorrect lenght of row number " + strconv.Itoa(counter) + ", all training rows must be at least 2")
		}
		name := flowerData[len(flowerData)-1]
		measurementsStrings := flowerData[0 : len(flowerData)-1]

		var measurements []float32

		for _, value := range measurementsStrings {
			measurementFloat, err := strconv.ParseFloat(strings.ReplaceAll(value, ",", "."), 32)
			if err != nil {
				return nil, errors.New("Unable to parse one of the measurements in line " + strconv.Itoa(counter) + " to float: " + err.Error())
			}
			measurements = append(measurements, float32(measurementFloat))
		}

		flower := Flower{
			name:         name,
			measurements: measurements,
		}

		flowers = append(flowers, flower)
	}

	return flowers, nil
}
