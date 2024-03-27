package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Flower struct {
	name         string
	measurements []float64
}

func main() {
	trainingPath := "./data/iris_training.txt"
	trainingFlowers := transformFlowers(flowerReader(trainingPath), "Iris-setosa")

	testPath := "./data/iris_test1.txt"
	testFlowers := transformFlowers(flowerReader(testPath), "Iris-setosa")

	w := createRandomArray(trainingFlowers)
	i := 0
	for i < 10 {
		for _, flower := range trainingFlowers {
			w = teachPerceptron(flower, w, 0.05)
		}
		i++
	}

	allCount := len(testFlowers)
	count := 0
	fmt.Println(w)

	for _, flower := range testFlowers {
		if flower.name == "1" && perceptron(flower.measurements, w) || flower.name == "0" && !perceptron(flower.measurements, w) {
			count += 1
		}
	}

	fmt.Println(count)
	fmt.Print(float64(count) / float64(allCount) * 100)
	fmt.Println("%")

	for true {

		var flowerString string
		fmt.Print("Type vector with white chars between dimentions and type of flower at the end (type quit to break): ")
		reader := bufio.NewReader(os.Stdin)
		flowerString, _ = reader.ReadString('\n')
		if flowerString == "quit" {
			break
		}
		flowerData := strings.Fields(flowerString)
		if len(flowerData) < 2 {
			fmt.Println("Incorrect size of input, every input has to have at least 1D vector and flower type")
			break
		}
		name := flowerData[len(flowerData)-1]
		measurementsStrings := flowerData[0 : len(flowerData)-1]

		var measurements []float64

		for _, value := range measurementsStrings {
			measurementFloat, err := strconv.ParseFloat(strings.ReplaceAll(value, ",", "."), 32)
			if err != nil {
				fmt.Println("Unable to parse " + value + " to float")
				break
			}
			measurements = append(measurements, float64(measurementFloat))
		}

		flower := Flower{
			name:         name,
			measurements: measurements,
		}
		predFlower := perceptron(flower.measurements, w)
		fmt.Print("Predicted flower: ")
		if predFlower {
			fmt.Println("Iris-setosa")
		} else {
			fmt.Println("Iris-nie-setosa")
		}

		if flower.name == "1" && perceptron(flower.measurements, w) || flower.name == "0" && !perceptron(flower.measurements, w) {
			fmt.Println("System predicted your flower")
		} else {
			fmt.Println("System didn't predict your flower")
		}
	}
}

func createRandomArray(flowers []Flower) []float64 {
	longestLength := 0

	for _, flower := range flowers {
		if len(flower.measurements) > longestLength {
			longestLength = len(flower.measurements)
		}
	}

	var array []float64
	i := 0
	for i < longestLength {
		array = append(array, rand.Float64()*2-1)
		i++
	}

	array = append(array, 1)

	return array
}

func perceptron(x []float64, w []float64) bool {
	sum := 0.0

	for i, weight := range w {
		if len(x) > i {
			sum += weight * x[i]
		}
	}

	return sum >= 0
}

func teachPerceptron(flower Flower, w []float64, alpha float64) []float64 {
	x := flower.measurements
	if !perceptron(x, w) && flower.name == "1" || perceptron(x, w) && flower.name == "0" {
		if flower.name == "0" {
			alpha *= -1
		}

		for i := range x {
			x[i] *= alpha
		}

		for i := range w {
			if i < len(x) {
				w[i] += x[i]
			}
		}
	}
	return w
}

func flowerReader(fileName string) []Flower {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()

	var flowers []Flower

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		flowerData := strings.Fields(line)
		name := flowerData[len(flowerData)-1]
		measurementsStrings := flowerData[:len(flowerData)-1]

		var measurements []float64

		for _, value := range measurementsStrings {
			measurementFloat, err := strconv.ParseFloat(strings.ReplaceAll(value, ",", "."), 64)
			if err != nil {
				fmt.Println("Error:", err)
			}
			measurements = append(measurements, measurementFloat)
		}

		flower := Flower{
			name:         name,
			measurements: measurements,
		}

		flowers = append(flowers, flower)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	return flowers
}

func transformFlowers(flowers []Flower, flowerName string) []Flower {
	for i := range flowers {
		if flowers[i].name != flowerName {
			flowers[i].name = "0"
		} else {
			flowers[i].name = "1"
		}

		flowers[i].measurements = append(flowers[i].measurements, -1)
	}

	return flowers
}
