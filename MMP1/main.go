package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Flower struct {
	name         string
	measurements []float32
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		flowerData := strings.Fields(line)
		if len(flowerData) < 2 {
			return nil, errors.New("Incorrect lenght of row, all training rows must be at least 2")
		}
		name := flowerData[len(flowerData)-1]
		measurementsStrings := flowerData[0 : len(flowerData)-1]

		var measurements []float32

		for _, value := range measurementsStrings {
			measurementFloat, err := strconv.ParseFloat(strings.ReplaceAll(value, ",", "."), 32)
			if err != nil {
				return nil, errors.New("Unable to parse one of the measurements to float: " + err.Error())
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
