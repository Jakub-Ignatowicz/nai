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
	var trainingPath string
	var testPath string

	fmt.Print("Provide path for training data(input \"default\" if you want to use default dir): ")
	fmt.Scan(&trainingPath)
	if trainingPath == "default" {
		trainingPath = "./data/iris_training.txt"
	}
	var trainingFlowers []Flower
	var err error

	trainingFlowers, err = flowerReader(trainingPath)

	for err != nil {
		fmt.Print("Incorrect path, try again. If you want to quit type \"quit\": ")
		fmt.Scan(&trainingPath)
		switch trainingPath {
		case "quit":
			return
		case "default":
			trainingPath = "./data/iris_training.txt"
			trainingFlowers, err = flowerReader(trainingPath)
		default:
			trainingFlowers, err = flowerReader(trainingPath)
		}
	}

	fmt.Print("Provide path for test data(input \"default\" if you want to use default dir): ")
	fmt.Scan(&testPath)
	if testPath == "default" {
		testPath = "./data/iris_test.txt"
	}
	var testFlowers []Flower

	testFlowers, err = flowerReader(testPath)

	for err != nil {
		fmt.Print("Incorrect path, try again. If you want to quit type \"quit\": ")
		fmt.Scan(&testPath)
		switch testPath {
		case "quit":
			return
		case "default":
			testPath = "./data/iris_test.txt"
			testFlowers, err = flowerReader(testPath)
		default:
			testFlowers, err = flowerReader(testPath)
		}
	}

	var kString string
	numberOfFlowers := len(trainingFlowers)
	formattedString := fmt.Sprintf("Provide number between 1 and %d for the number of closest neighbors tested: ", numberOfFlowers)
	fmt.Print(formattedString)
	fmt.Scan(&kString)
	if kString == "quit" {
		return
	}

	var k int
	k, err = strconv.Atoi(kString)

	for err != nil || k < 1 || k > numberOfFlowers {
		formattedString = fmt.Sprintf("Incorrect k, please provide whole number between 1 and %d(type \"quit\" if you want to stop a program): ", numberOfFlowers)
		fmt.Print(formattedString)
		fmt.Scan(&kString)
		if kString == "quit" {
			return
		}
		k, err = strconv.Atoi(kString)
	}

	err = predictFlowers(trainingFlowers, testFlowers, k)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println("Here is set of commands you can use:")
	fmt.Println(" help - print all possible commands again")
	fmt.Println(" setTraining - lets you change path to training data")
	fmt.Println(" setTest - lets you change path to test data")
	fmt.Println(" setK - lets you change k neighbors tested")
	fmt.Println(" testFlower - lets you input flower for testing")
	fmt.Println(" testFlowers - test your test set on training set")
	fmt.Println(" quit - leaves the program")

	var command string

	for true {

		fmt.Print("> ")
		fmt.Scan(&command)
		switch command {
		case "quit":
			return
		case "help":
			fmt.Println("Here is set of commands you can use:")
			fmt.Println(" help - print all possible commands again")
			fmt.Println(" setTraining - lets you change path to training data")
			fmt.Println(" setTest - lets you change path to test data")
			fmt.Println(" setK - lets you change k neighbors tested")
			fmt.Println(" testFlower - lets you input flower for testing")
			fmt.Println(" testFlowers - test your test set on training set")
			fmt.Println(" quit - leaves the program")
		case "setTraining":
			fmt.Print("Provide path or type \"default\": ")
			fmt.Scan(&trainingPath)
			if trainingPath == "default" {
				trainingPath = "./data/iris_training.txt"
			}
			trainingFlowers, err = flowerReader(trainingPath)
			if err != nil {
				fmt.Println("Couldn't read this path, try again")
				err = nil
			} else {
				fmt.Println("Read flowers from " + trainingPath)
			}
		case "setTest":
			fmt.Print("Provide path or type \"default\": ")
			fmt.Scan(&testPath)
			if testPath == "default" {
				testPath = "./data/iris_test.txt"
			}
			testFlowers, err = flowerReader(testPath)
			if err != nil {
				fmt.Println("Couldn't read this path, try again")
				err = nil
			} else {
				fmt.Println("Read flowers from " + testPath)
			}
		case "testFlower":
			var flowerString string
			fmt.Print("Type vector with white chars between dimentions and type of flower at the end: ")
			reader := bufio.NewReader(os.Stdin)
			flowerString, _ = reader.ReadString('\n')
			flowerData := strings.Fields(flowerString)
			if len(flowerData) < 2 {
				fmt.Println("Incorrect size of input, every input has to have at least 1D vector and flower type")
				break
			}
			name := flowerData[len(flowerData)-1]
			measurementsStrings := flowerData[0 : len(flowerData)-1]

			var measurements []float32

			for _, value := range measurementsStrings {
				measurementFloat, err := strconv.ParseFloat(strings.ReplaceAll(value, ",", "."), 32)
				if err != nil {
					fmt.Println("Unable to parse " + value + " to float")
					break
				}
				measurements = append(measurements, float32(measurementFloat))
			}

			flower := Flower{
				name:         name,
				measurements: measurements,
			}
			predFlower, err := predictFlower(trainingFlowers, flower, k)
			if err != nil {
				fmt.Println("Error why predicting flower: " + err.Error())
			}

			if predFlower {
				fmt.Println("System predicted your flower")
			} else {
				fmt.Println("System didn't predict your flower")
			}
		case "testFlowers":
			predictFlowers(trainingFlowers, testFlowers, k)
		case "setK":
			fmt.Print("Provide k: ")
			fmt.Scan(&kString)
			k, err = strconv.Atoi(kString)
			if err != nil || k > len(trainingFlowers) || k < 1 {
				formattedString = fmt.Sprintf("Incorrect k, please try again and provide whole number between 1 and %d", numberOfFlowers)
				fmt.Println(formattedString)
				err = nil
			} else {
				fmt.Println("New k has been set")
			}
		default:
			fmt.Println("Incorrect command, try again(for list of command type \"help\")")
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

func predictFlower(trainFlowers []Flower, testFlower Flower, k int) (bool, error) {
	if k > len(trainFlowers) || k < 1 {
		return false, errors.New("k must be greater than 0 and smaller than training dataset size")
	}

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

	return predictedFlower == testFlower.name, nil
}

func predictFlowers(trainFlowers []Flower, testFlowers []Flower, k int) error {
	if k > len(trainFlowers) || k < 1 {
		return errors.New("k must be greater than 0 and smaller than training dataset size")
	}

	testCount := float64(len(testFlowers))
	correctCount := 0.0

	for _, testFlower := range testFlowers {
		correctPred, err := predictFlower(trainFlowers, testFlower, k)
		if err != nil {
			return err
		}
		if correctPred {
			correctCount++
		}
	}
	percentage := (correctCount / testCount) * 100

	fmt.Printf("For provided datasets prediction was correct %.1f%% of the time\n", percentage)
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
