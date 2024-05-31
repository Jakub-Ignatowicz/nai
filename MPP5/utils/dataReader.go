package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func DataReader(filePath string) ([][]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ans [][]float64

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Replace(scanner.Text(), ",", ".", -1)

		numberStrings := strings.Fields(line)

		var numbers []float64

		for _, numberString := range numberStrings {
			number, err := strconv.ParseFloat(numberString, 64)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, number)
		}

		ans = append(ans, numbers)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ans, nil
}
