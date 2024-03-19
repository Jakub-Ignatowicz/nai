package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Flower struct {
	name         string
	measurements []float32
}

type FlowerDistance struct {
	flower   Flower
	distance float32
}

func (f1 Flower) Distance(f2 Flower) FlowerDistance {
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

	return FlowerDistance{
		flower:   f2,
		distance: float32(math.Sqrt(distance)),
	}
}

func main() {
	for k := 0; k < 120; k += 5 {
		err := testFlower("./data/iris_training.txt", "./data/iris_test.txt", k)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func flowerReader(fileName string) ([]Flower, error) {
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

func testFlower(trainingFileName string, testFileName string, k int) error {
	trainFlowers, e1 := flowerReader(trainingFileName)
	if e1 != nil {
		return e1
	}
	if k > len(trainFlowers) || k < 1 {
		return errors.New("k must be greater than 0 and smaller than training dataset size")
	}
	testFlowers, e2 := flowerReader(testFileName)
	if e2 != nil {
		return e2
	}

	testCount := float64(len(testFlowers))
	correctCount := 0.0

	for _, testFlower := range testFlowers {
		var flowerDistances []FlowerDistance
		for _, trainFlower := range trainFlowers {
			flowerDistances = append(flowerDistances, testFlower.Distance(trainFlower))
		}

		sort.Slice(flowerDistances, func(i, j int) bool {
			return flowerDistances[i].distance < flowerDistances[j].distance
		})

		namesCount := make(map[string]int)

		for i := 0; i < k; i++ {
			currName := flowerDistances[i].flower.name
			value, exists := namesCount[currName]
			if exists {
				namesCount[currName] = value + 1
			} else {
				namesCount[currName] = 1
			}
		}

		predictedFlower := maxKeyInMap(namesCount)

		if predictedFlower == testFlower.name {
			correctCount++
		}
	}
	percentage := (correctCount / testCount) * 100

	fmt.Println(k)
	fmt.Printf("Prediction was correct %.1f%% of the time\n", percentage)
	return nil
}

func maxKeyInMap(m map[string]int) string {
	maxKey := ""
	maxValue := math.MinInt

	for key, value := range m {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}
	return maxKey
}
