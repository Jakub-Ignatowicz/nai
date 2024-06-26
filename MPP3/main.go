package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"mpp3/utils"
	"os"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	files, err := utils.DataReader("./data")

	if err != nil {
		println(err.Error())
	}

	languageWeights := makeLanguageWeightsMap(files)

	accuracy := 0.0

	count := 0

	for accuracy < 0.99 {
		for _, file := range files {
			for language, weights := range languageWeights {
				var expectedValue float64 = -1

				if file.Language == language {
					expectedValue = 1
				}

				newW, err := utils.TeachPerceptron(expectedValue, file.ProportionVector, weights, 0.01)

				if err != nil {
					println(err)
					return
				}

				languageWeights[language] = newW

			}
		}
		accuracy = testPerceptrons(languageWeights, files)
		print("Current accuracy is equal to ")
		fmt.Printf("%.2f\n", accuracy)
		count++
	}
	print(count)
	println(" epok")

	for true {
		fmt.Print("Provide text here(typq 'q' to quit): ")
		scanner := bufio.NewScanner(os.Stdout)
		if scanner.Scan() {
			sentence := scanner.Text()
			if sentence == "q" {
				return
			}

			fmt.Print("Predicted language: ")
			fmt.Println(testLangauge(languageWeights, sentence))
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading standard input:", err)
		}
	}
}

func testPerceptrons(languageWeights map[string][]float64, files []utils.File) float64 {
	correctGuesses := 0.0
	for _, file := range files {
		maxWeight := math.SmallestNonzeroFloat64
		currLang := ""
		for language, weights := range languageWeights {
			perceptron, err := utils.Perceptron(file.ProportionVector, weights)
			if err != nil {
				println(err)
			} else {
				if maxWeight < perceptron {
					maxWeight = perceptron
					currLang = language
				}
			}
		}

		if file.Language == currLang {
			correctGuesses++
		}
	}

	return correctGuesses / float64(len(files))
}

func testLangauge(languageWeights map[string][]float64, text string) string {
	maxWeight := math.SmallestNonzeroFloat64
	currLang := ""
	for language, weights := range languageWeights {
		perceptron, err := utils.Perceptron(utils.Normalize(utils.CountAllLetters(text)), weights)
		if err != nil {
			fmt.Println(err)
		} else {
			if maxWeight < perceptron {
				maxWeight = perceptron
				currLang = language
			}
		}
	}
	return currLang
}

func randomWeights(length int) []float64 {
	weights := make([]float64, length)
	for i := 0; i < length; i++ {
		weights[i] = rand.Float64()
	}
	return weights
}

func makeLanguageWeightsMap(files []utils.File) map[string][]float64 {
	languageSet := mapset.NewSet[string]()

	for _, file := range files {
		languageSet.Add(file.Language)
	}

	languageWeights := make(map[string][]float64)

	for lang := range languageSet.Iterator().C {
		languageWeights[lang] = randomWeights(26)
	}

	return languageWeights
}
