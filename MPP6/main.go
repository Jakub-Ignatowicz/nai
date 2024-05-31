package main

import (
	"container/heap"
	"fmt"
	"math"
	"mpp6/utils"
)

func countLetters(sentence string) map[rune]int {
	letterCount := make(map[rune]int)

	for _, char := range sentence {
		letterCount[char]++
	}

	return letterCount
}

func createInitialPriorityQueue(sentence string) utils.PriorityQueue {
	letterCounts := countLetters(sentence)

	pq := make(utils.PriorityQueue, len(letterCounts))

	i := 0
	for key, value := range letterCounts {
		pq[i] = utils.Item{
			Name:   string(key),
			Weight: value,
		}
		i++
	}
	heap.Init(&pq)

	return pq
}

func main() {
	str := "hakuma tata ma"
	pq := createInitialPriorityQueue(str)

	bitSizeOfChar := 1

	for math.Pow(2, float64(bitSizeOfChar)) < float64(pq.Len()) {
		bitSizeOfChar++
	}

	result := fmt.Sprintf("Total bit length of not encoded string is: %d", bitSizeOfChar*len(str))

	fmt.Println(result)

	encodedChars := make(map[string]string)

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(utils.Item)
		right := heap.Pop(&pq).(utils.Item)

		newName := left.Name + right.Name
		newWeight := left.Weight + right.Weight

		newItem := utils.Item{
			Name:   newName,
			Weight: newWeight,
		}

		heap.Push(&pq, newItem)

		for _, char := range left.Name {
			encodedChars[string(char)] = "0" + encodedChars[string(char)]
		}

		for _, char := range right.Name {
			encodedChars[string(char)] = "1" + encodedChars[string(char)]
		}
	}
}
